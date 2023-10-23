package handler

import (
	"app4/database"
	pb "app4/proto"
	"app4/service"
	"context"
	_ "github.com/mwitkow/go-proto-validators"
	"google.golang.org/grpc"
)

type ServerService struct {
	UserService service.UserServiceInterface
	GrpcServer  *grpc.Server
	pb.UnimplementedUserServiceServer
}

func NewServerService(UserService service.UserServiceInterface) *ServerService {
	userGrpc := &ServerService{UserService: UserService}
	return userGrpc
}

func (ss *ServerService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	//
	//err := req.Validate()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return nil, nil
	//
	data, err := ss.transformUserRPC(req)
	if err != nil {
		return nil, err
	}
	//	fmt.Println(*data)

	userID, err := ss.UserService.Create(*data)
	if err != nil {
		return nil, err
	}

	// no need to check error as transformUserModel() used only with previously created users
	user, _ := ss.UserService.Get(userID)
	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	ID := req.GetID()

	user, err := ss.UserService.Get(ID)
	if err != nil {
		return nil, err
	}

	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Update(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data, err := ss.transformUserRPC(req)
	if err != nil {
		return nil, err
	}
	//	data := ss.transformUserRPC(req)
	ID := req.GetID()
	//	fmt.Println(*data)

	updatedID, err := ss.UserService.Update(ID, *data)
	if err != nil {
		return nil, err
	}

	user, _ := ss.UserService.Get(updatedID)

	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Delete(ctx context.Context, req *pb.SingleUserRequest) (*pb.SuccessResponse, error) {
	ID := req.GetID()

	err := ss.UserService.Delete(ID)
	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Response: "User Successfully deleted"}, nil
}

func (ss *ServerService) transformUserRPC(req *pb.CreateUserRequest) (*database.User, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return nil, nil
	//return &database.User{
	//	//		ID:        req.GetID(),
	//	Nickname:  req.GetNickname(),
	//	FirstName: req.GetFirstName(),
	//	LastName:  req.GetLastName(),
	//	Password:  req.GetPassword(),
	//	//CreatedAt: req.GetCreatedAt(),
	//	//		UpdatedAt: req.GetUpdatedAt(),
	//	//		DeletedAt: req.GetDeletedAt(),
	//}, nil
}

func (ss *ServerService) transformUserModel(user database.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		//		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
