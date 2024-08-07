package trade

import (
	"context"

	pb "example.com/grpc-todo/gen"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTradeServiceServer
	service   TradeService
	validator *protovalidate.Validator
}

func NewTradeGrpcServer() *server {
	validator, _ := protovalidate.New()
	return &server{service: NewTradeCacheService(), validator: validator}
}

func (s *server) Put(ctx context.Context, in *pb.Trade) (*pb.Empty, error) {
	if err := s.validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "userID not found")
	}

	return &pb.Empty{}, s.service.Put(ctx, userID, TradeAction{
		Symbol: in.Symbol.Enum().String(),
		Amount: in.Amount,
	})
}

func (s *server) Call(ctx context.Context, in *pb.Trade) (*pb.Empty, error) {
	if err := s.validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "userID not found")
	}

	return &pb.Empty{}, s.service.Call(ctx, userID, TradeAction{
		Symbol: in.Symbol.Enum().String(),
		Amount: in.Amount,
	})
}

func (s *server) TakeProfit(ctx context.Context, in *pb.Trade) (*pb.Empty, error) {
	if err := s.validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// userID, ok := ctx.Value("userID").(int)
	// if !ok {
	// 	return nil, status.Error(codes.InvalidArgument, "userID not found")
	// }

	return &pb.Empty{}, s.service.TakeProfit(ctx, 1, TradeAction{
		Symbol: in.Symbol.Enum().String(),
		Amount: in.Amount,
	})
}

func (s *server) StopLose(ctx context.Context, in *pb.Trade) (*pb.Empty, error) {
	if err := s.validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// userID, ok := ctx.Value("userID").(int)
	// if !ok {
	// 	return nil, status.Error(codes.InvalidArgument, "userID not found")
	// }

	return &pb.Empty{}, s.service.StopLose(ctx, 1, TradeAction{
		Symbol: in.Symbol.Enum().String(),
		Amount: in.Amount,
	})
}
