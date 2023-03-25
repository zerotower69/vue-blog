package models

// CommentModel 是文章的评论 是多对一的关系
type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` //父级评论
	ParentCommentID    uint            `json:"parent_comment_id"`                               //父评论id

	Content      string       `json:"content"`                               //评论内容
	LikeCount    int          `gorm:"size:8;default:0" json:"like_count"`    //点赞数
	CommentCount int          `gorm:"size:8;default:0" json:"comment_count"` //子评论数
	Article      ArticleModel `gorm:"foreignKey:ArticleID" json:"article"`   //关联的文章
	ArticleID    uint         `json:"article_id"`                            //文章id
	User         UserModel    `json:"user"`                                  //关联的用户
	UserID       uint         `json:"user_id"`                               //评论的用户
}
