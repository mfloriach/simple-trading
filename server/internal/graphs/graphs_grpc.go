package graphs

import (
	"context"
	"log"

	pb "example.com/grpc-todo/gen"
	"example.com/grpc-todo/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedGraphServiceServer
	validator  utils.Validator
	pool       ApiPool
	repository GraphRepository
}

func NewGrpcServer(validator utils.Validator, pool ApiPool) *server {
	return &server{
		validator:  validator,
		pool:       pool,
		repository: NewTimescaleRepo(),
	}
}

func (s *server) GetBySymbol(ctx context.Context, req *pb.Symbol) (*pb.Candles, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	candles, err := s.repository.GetBySymbol(ctx, req.Id.String(), TimeRange(*req.TimeRange.Enum()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res []*pb.Candle
	for _, candle := range candles {
		res = append(res, &pb.Candle{
			Symbol:    req.Id,
			High:      uint32(candle.High),
			Open:      uint32(candle.Open),
			Close:     uint32(candle.Close),
			Low:       uint32(candle.Low),
			TimeRange: req.TimeRange,
			CreatedAt: timestamppb.Now(),
		})
	}

	return &pb.Candles{Candles: res}, nil

}

func (s *server) GetBySymbolStream(req *pb.Symbol, srv pb.GraphService_GetBySymbolStreamServer) error {
	if err := s.validator.Validate(req); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	api := s.pool.Get(req.Id.String())
	if api == nil {
		return status.Error(codes.Internal, "no available apis")
	}

	position, ch := api.Subscribe()

	for {
		candle := <-ch

		if err := srv.SendMsg(&pb.Candle{
			Symbol:    req.Id,
			High:      uint32(candle.High),
			Open:      uint32(candle.Open),
			Close:     uint32(candle.Close),
			Low:       uint32(candle.Low),
			TimeRange: req.TimeRange,
			CreatedAt: timestamppb.Now(),
		}); err != nil {
			log.Println("error generating response")
			api.UnSubscribe(position)
			return err
		}
	}
}
