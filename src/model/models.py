from datetime import datetime
from re import S
from peewee import *

from config import env


class BaseModel(Model):
    class Meta:
        database = env.DB


class User(BaseModel):
    created_at = DateTimeField(constraints=[SQL("DEFAULT CURRENT_TIMESTAMP")])
    updated_at = DateTimeField(
        constraints=[SQL("DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")]
    )
    deleted_at = DateTimeField(null=True)
    user_name = CharField(max_length=40, unique=True, null=False)
    password = CharField(max_length=256, null=False)
    following_count = IntegerField(default=0)
    follower_count = IntegerField(default=0)
    role = IntegerField(default=1)
    avatar = CharField(max_length=255)
    background_image = CharField(max_length=255)
    signature = CharField(max_length=120)


class Video(BaseModel):
    created_at = DateTimeField(constraints=[SQL("DEFAULT CURRENT_TIMESTAMP")])
    updated_at = DateTimeField(
        constraints=[SQL("DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")]
    )
    deleted_at = DateTimeField(null=True)
    update_time = DateTimeField(default=datetime.now, null=False, index=True)
    author_id = ForeignKeyField(User, backref="videos", index=True)
    play_url = CharField(max_length=255, null=False)
    cover_url = CharField(max_length=255)
    favorite_count = IntegerField(default=0)
    comment_count = IntegerField(default=0)
    title = CharField(max_length=50, null=False)


class Comment(BaseModel):
    created_at = DateTimeField(constraints=[SQL("DEFAULT CURRENT_TIMESTAMP")])
    updated_at = DateTimeField(
        constraints=[SQL("DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")]
    )
    deleted_at = DateTimeField(null=True)
    video_id = ForeignKeyField(Video, backref="comments", index=True)
    user_id = ForeignKeyField(User, backref="comments", index=True)
    content = CharField(max_length=255, null=False)


class UserFavoriteVideos(BaseModel):
    created_at = DateTimeField(constraints=[SQL("DEFAULT CURRENT_TIMESTAMP")])
    updated_at = DateTimeField(
        constraints=[SQL("DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")]
    )
    deleted_at = DateTimeField(null=True)
    user_id = ForeignKeyField(User, backref="favorite_videos")
    video_id = ForeignKeyField(Video, backref="favorited_by_users")

    class Meta:
        indexes = (
            # 您可以在这里创建复合索引，确保用户和视频的组合是唯一的
            (("user_id", "video_id"), True),
        )


class Relation(BaseModel):
    created_at = DateTimeField(constraints=[SQL("DEFAULT CURRENT_TIMESTAMP")])
    updated_at = DateTimeField(
        constraints=[SQL("DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")]
    )
    deleted_at = DateTimeField(null=True)
    user_id = ForeignKeyField(
        User, backref="followings", index=True, unique=True, null=False
    )
    to_user_id = ForeignKeyField(
        User, backref="followers", index=True, unique=True, null=False
    )


if __name__ == "__main__":
    env.DB.create_tables([User, Video, Comment, UserFavoriteVideos, Relation])
