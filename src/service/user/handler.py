import grpc
from loguru import logger

import peewee

from model.models import User
from rpc.user import user_pb2, user_pb2_grpc


class UserServicer(user_pb2_grpc.UserServiceServicer):
    @logger.catch
    def GetUserInfo(self, request, context):
        resp = user_pb2.UserResponse()

        userId = request.user_id
        try:
            user = User.get_by_id(userId)

        except peewee.DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("User not found")
            resp.status_code = 50021
            resp.status_msg = "User not found"
            return resp

        resp.status_code = 1
        resp.status_msg = "success"
        resp.user = user_pb2.User(id=user.id, name=user.name, is_follow=True)

        return resp

    @logger.catch
    def GetUserExistInformation(self, request, context):
        pass
