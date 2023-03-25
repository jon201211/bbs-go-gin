package controller

import (
	"bbs-go/controller/admin"
	"bbs-go/controller/api"
	"bbs-go/middleware"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func Router(app *gin.Engine) {

	//	app.Use(middleware.BodyLogMiddleware)

	// api
	api_router := app.Group("/api")

	topicController := &api.TopicController{}
	api_router.GET("/topic/nodes", topicController.GetNodes)
	api_router.GET("/topic/node", topicController.GetNode)
	api_router.POST("/topic/create", topicController.PostCreate)
	api_router.GET("/topic/edit/:id", topicController.GetEditBy)
	api_router.POST("/topic/edit/:id", topicController.PostEditBy)
	api_router.POST("/topic/delete/:id", topicController.PostDeleteBy)
	api_router.POST("/topic/recommend/:id", topicController.PostRecommendBy)
	api_router.GET("/topic/:id", topicController.GetBy)
	api_router.POST("/topic/like/:id", topicController.PostLikeBy)
	api_router.GET("/topic/recentlikes/:id", topicController.GetRecentlikesBy)
	api_router.GET("/topic/recent", topicController.GetRecent)
	api_router.GET("/topic/user/topics", topicController.GetUserTopics)
	api_router.GET("/topic/topics", topicController.GetTopics)
	api_router.GET("/topic/tag/topics", topicController.GetTagTopics)
	api_router.GET("/topic/favorite/:id", topicController.GetFavoriteBy)
	api_router.GET("/topic/recommend", topicController.GetRecommend)
	api_router.GET("/topic/newest", topicController.GetNewest)

	//api_router.Party("/article").Handle(new(api.ArticleController))
	articleController := &api.ArticleController{}
	api_router.GET("/article/:id", articleController.GetBy)
	api_router.POST("/article/create", articleController.PostCreate)
	api_router.GET("/article/edit/:id", articleController.GetEditBy)
	api_router.POST("/article/edit/:id", articleController.PostEditBy)
	api_router.POST("/article/delete/:id", articleController.PostDeleteBy)
	api_router.POST("/article/favorite/:id", articleController.PostFavoriteBy)
	api_router.GET("/article/redirect/:id", articleController.GetRedirectBy)
	api_router.GET("/article/recent", articleController.GetRecent)
	api_router.GET("/article/user/articles", articleController.GetUserArticles)
	api_router.GET("/article/articles", articleController.GetArticles)
	api_router.GET("/article/tag/articles", articleController.GetTagArticles)
	api_router.GET("/article/user/newest/:userId", articleController.GetUserNewestBy)
	api_router.GET("/article/nearly/:id", articleController.GetNearlyBy)
	api_router.GET("/article/related/:id", articleController.GetRelatedBy)
	api_router.GET("/article/recommend", articleController.GetRecommend)
	api_router.GET("/article/newest", articleController.GetNewest)
	api_router.GET("/article/hot", articleController.GetHot)

	//api_router.Party("/project").Handle(new(api.ProjectController))
	projectController := &api.ProjectController{}
	api_router.GET("/project/:id", projectController.GetBy)
	api_router.GET("/project/projects", projectController.GetProjects)

	//api_router.Party("/login").Handle(new(api.LoginController))
	loginController := &api.LoginController{}
	api_router.POST("/login/signin", loginController.PostSignin)
	api_router.POST("/login/signup", loginController.PostSignup)
	api_router.GET("/login/signout", loginController.GetSignout)

	//api_router.Party("/user").Handle(new(api.UserController))
	userController := &api.UserController{}
	api_router.GET("/user/current", userController.GetCurrent)
	api_router.GET("/user/:id", userController.GetBy)
	api_router.POST("/user/edit/:id", userController.PostEditBy)
	api_router.POST("/user/update/avatar", userController.PostUpdateAvatar)
	api_router.POST("/user/set/username", userController.PostSetUsername)
	api_router.POST("/user/set/email", userController.PostSetEmail)
	api_router.POST("/user/set/password", userController.PostSetPassword)
	api_router.POST("/user/update/password", userController.PostUpdatePassword)
	api_router.POST("/user/set/background/image", userController.PostSetBackgroundImage)
	api_router.GET("/user/favorites", userController.GetFavorites)
	api_router.GET("/user/msgrecent", userController.GetMsgrecent)
	api_router.GET("/user/messages", userController.GetMessages)
	api_router.GET("/user/scorelogs", userController.GetScorelogs)
	api_router.GET("/user/score/rank", userController.GetScoreRank)
	api_router.POST("/user/forbidden", userController.PostForbidden)
	api_router.GET("/user/email/verifiy", userController.GetEmailVerify)
	api_router.POST("/user/email/verifiy", userController.PostEmailVerify)

	//api_router.Party("/tag").Handle(new(api.TagController))
	tagController := &api.TagController{}
	api_router.GET("/tag/:id", tagController.GetBy)
	api_router.GET("/tag/tags", tagController.GetTags)
	api_router.POST("/tag/autocomplete", tagController.PostAutocomplete)

	//api_router.Party("/comment").Handle(new(api.CommentController))
	commentController := &api.CommentController{}
	api_router.GET("/comment/fuck", commentController.GetFuck)
	api_router.POST("/comment/create", commentController.PostCreate)
	api_router.GET("/comment/list", commentController.GetList)

	//api_router.Party("/favorite").Handle(new(api.FavoriteController))
	favoriteController := &api.FavoriteController{}
	api_router.GET("/favorite/favorited", favoriteController.GetFavorited)
	api_router.GET("/favorite/delete", favoriteController.GetDelete)

	//api_router.Party("/like").Handle(new(api.LikeController))
	likeController := &api.LikeController{}
	api_router.GET("/like/is/liked", likeController.GetIsLiked)
	api_router.GET("/like/liked", likeController.GetLiked)

	//api_router.Party("/checkin").Handle(new(api.CheckinController))
	checkinController := &api.CheckinController{}
	api_router.POST("/checkin/checkin", checkinController.PostCheckin)
	api_router.GET("/checkin/checkin", checkinController.GetCheckin)
	api_router.GET("/checkin/rank", checkinController.GetRank)

	//api_router.Party("/config").Handle(new(api.ConfigController))
	configController := &api.ConfigController{}
	api_router.GET("/config/configs", configController.GetConfigs)

	//api_router.Party("/upload").Handle(new(api.UploadController))
	uploadController := &api.UploadController{}
	api_router.POST("/upload", uploadController.Post)
	api_router.POST("/upload/editor", uploadController.PostEditor)
	api_router.POST("/upload/fetch", uploadController.PostFetch)

	//api_router.Party("/link").Handle(new(api.LinkController))
	linkController := &api.LinkController{}
	api_router.GET("/link/:id", linkController.GetBy)
	api_router.GET("/link/links", linkController.GetLinks)
	api_router.GET("/link/toplinks", linkController.GetToplinks)

	//api_router.Party("/captcha").Handle(new(api.CaptchaController))
	captchaController := &api.CaptchaController{}
	api_router.GET("/captcha/request", captchaController.GetRequest)
	api_router.GET("/captcha/show", captchaController.GetShow)
	api_router.GET("/captcha/verify", captchaController.GetVerify)

	//api_router.Party("/qq/login").Handle(new(api.QQLoginController))
	qqLoginController := &api.QQLoginController{}
	api_router.GET("/qq/login/authorize", qqLoginController.GetAuthorize)
	api_router.GET("/qq/login/callback", qqLoginController.GetCallback)

	//api_router.Party("/github/login").Handle(new(api.GithubLoginController))
	githubLoginController := &api.GithubLoginController{}
	api_router.GET("/github/login/authorize", githubLoginController.GetAuthorize)
	api_router.GET("/github/login/callback", githubLoginController.GetCallback)

	//api_router.Party("/osc/login").Handle(new(api.OscLoginController))
	oscLoginController := &api.OscLoginController{}
	api_router.GET("/osc/login/authorize", oscLoginController.GetAuthorize)
	api_router.GET("/osc/login/callback", oscLoginController.GetCallback)

	//api_router.Party("/search").Handle(new(api.SearchController))
	searchController := &api.SearchController{}
	api_router.GET("/search/any/reindex", searchController.AnyReindex)
	api_router.POST("/search/topic", searchController.PostTopic)

	//api_router.Party("/spider").Handle(new(api.SpiderController))
	spiderController := &api.SpiderController{}
	api_router.POST("/spider/wx/publish", spiderController.PostWxPublish)
	api_router.POST("/spider/article/publish", spiderController.PostArticlePublish)
	api_router.POST("/spider/comment/publish", spiderController.PostCommentPublish)
	api_router.POST("/spider/project/publish", spiderController.PostProjectPublish)

	//api_router.Party("/fans").Handle(new(api.FansController))
	fansController := &api.FansController{}
	api_router.POST("/fans/follow", fansController.PostFollow)
	api_router.POST("/fans/unfollow", fansController.PostUnfollow)
	api_router.GET("/fans/isfollowed", fansController.GetIsfollowed)
	api_router.GET("/fans/fans", fansController.GetFans)
	api_router.GET("/fans/follows", fansController.GetFollows)
	api_router.GET("/fans/recent/fans", fansController.GetRecentFans)
	api_router.GET("/fans/recent/follow", fansController.GetRecentFollow)

	//api_router.Party("/feed").Handle(new(api.FeedController))
	feedController := &api.FeedController{}
	api_router.GET("/feed/topics", feedController.GetTopics)

	// admin
	admin_router := app.Group("/api/admin")

	admin_router.Use(middleware.AdminAuth)

	//admin_router.Party("/common").Handle(new(admin.CommonController))
	commonController := &admin.CommonController{}
	admin_router.GET("/common/systeminfo", commonController.GetSysteminfo)

	//admin_router.Party("/user").Handle(new(admin.UserController))
	adminUserController := &admin.UserController{}
	admin_router.GET("/user/systeminfo", adminUserController.GetSynccount)
	admin_router.Any("/user/list", adminUserController.AnyList)
	admin_router.GET("/user/:id", adminUserController.GetBy)
	admin_router.POST("/user/create", adminUserController.PostCreate)
	admin_router.POST("/user/update", adminUserController.PostUpdate)
	admin_router.POST("/user/forbidden", adminUserController.PostForbidden)

	//admin_router.Party("/third-account").Handle(new(admin.ThirdAccountController))
	thirdAccountController := &admin.ThirdAccountController{}
	admin_router.GET("/third-account/:id", thirdAccountController.GetBy)
	admin_router.Any("/third-account/list", thirdAccountController.AnyList)
	admin_router.POST("/third-account/create", thirdAccountController.PostCreate)
	admin_router.POST("/third-account/update", thirdAccountController.PostUpdate)

	//admin_router.Party("/tag").Handle(new(admin.TagController))
	adminTagController := &admin.TagController{}
	admin_router.GET("/tag/:id", adminTagController.GetBy)
	admin_router.Any("/tag/list", adminTagController.AnyList)
	admin_router.POST("/tag/create", adminTagController.PostCreate)
	admin_router.POST("/tag/update", adminTagController.PostUpdate)
	admin_router.GET("/tag/autocomplete", adminTagController.GetAutocomplete)
	admin_router.GET("/tag/tags", adminTagController.GetTags)

	//admin_router.Party("/article").Handle(new(admin.ArticleController))
	adminArticleController := &admin.ArticleController{}
	admin_router.GET("/article/sitemap", adminArticleController.GetSitemap)
	admin_router.GET("/article/:id", adminArticleController.GetBy)
	admin_router.Any("/article/list", adminArticleController.AnyList)
	admin_router.GET("/article/recent", adminArticleController.GetRecent)
	admin_router.POST("/article/update", adminArticleController.PostUpdate)
	admin_router.GET("/article/tags", adminArticleController.GetTags)
	admin_router.PUT("/article/tags", adminArticleController.PutTags)
	admin_router.POST("/article/delete", adminArticleController.PostDelete)
	admin_router.POST("/article/pending", adminArticleController.PostPending)

	//admin_router.Party("/comment").Handle(new(admin.CommentController))
	adminCommentController := &admin.CommentController{}
	admin_router.GET("/comment/:id", adminCommentController.GetBy)
	admin_router.Any("/comment/list", adminCommentController.AnyList)
	admin_router.POST("/comment/delete/:id", adminCommentController.PostDeleteBy)

	//admin_router.Party("/favorite").Handle(new(admin.FavoriteController))
	adminFavoriteController := &admin.FavoriteController{}
	admin_router.GET("/favorite/:id", adminFavoriteController.GetBy)
	admin_router.Any("/favorite/list", adminFavoriteController.AnyList)
	admin_router.POST("/favorite/create", adminFavoriteController.PostCreate)
	admin_router.POST("/favorite/update", adminFavoriteController.PostUpdate)

	//admin_router.Party("/article-tag").Handle(new(admin.ArticleTagController))
	adminArticleTagController := &admin.ArticleTagController{}
	admin_router.GET("/article-tag/:id", adminArticleTagController.GetBy)
	admin_router.Any("/article-tag/list", adminArticleTagController.AnyList)
	admin_router.POST("/article-tag/create", adminArticleTagController.PostCreate)
	admin_router.POST("/article-tag/update", adminArticleTagController.PostUpdate)

	//admin_router.Party("/topic").Handle(new(admin.TopicController))
	adminTopicController := &admin.TopicController{}
	admin_router.GET("/topic/:id", adminTopicController.GetBy)
	admin_router.Any("/topic/list", adminTopicController.AnyList)
	admin_router.POST("/topic/recommend", adminTopicController.PostRecommend)
	admin_router.DELETE("/topic/recommend", adminTopicController.DeleteRecommend)
	admin_router.POST("/topic/delete", adminTopicController.PostDelete)
	admin_router.POST("/topic/undelete", adminTopicController.PostUndelete)

	//admin_router.Party("/topic-node").Handle(new(admin.TopicNodeController))
	adminTopicNodeController := &admin.TopicNodeController{}
	admin_router.GET("/topic-node/:id", adminTopicNodeController.GetBy)
	admin_router.Any("/topic-node/list", adminTopicNodeController.AnyList)
	admin_router.POST("/topic-node/create", adminTopicNodeController.PostCreate)
	admin_router.POST("/topic-node/update", adminTopicNodeController.PostUpdate)
	admin_router.GET("/topic-node/nodes", adminTopicNodeController.GetNodes)

	//admin_router.Party("/sys-config").Handle(new(admin.SysConfigController))
	adminSysConfigController := &admin.SysConfigController{}
	admin_router.GET("/sys-config/:id", adminSysConfigController.GetBy)
	admin_router.Any("/sys-config/list", adminSysConfigController.AnyList)
	admin_router.GET("/sys-config/all", adminSysConfigController.GetAll)
	admin_router.POST("/sys-config/save", adminSysConfigController.PostSave)

	//admin_router.Party("/link").Handle(new(admin.LinkController))
	adminLinkController := &admin.LinkController{}
	admin_router.GET("/link/:id", adminLinkController.GetBy)
	admin_router.Any("/link/list", adminLinkController.AnyList)
	admin_router.POST("/link/create", adminLinkController.PostCreate)
	admin_router.POST("/link/update", adminLinkController.PostUpdate)
	admin_router.GET("/link/detect", adminLinkController.GetDetect)

	//admin_router.Party("/user-score-log").Handle(new(admin.UserScoreLogController))
	adminUserScoreLogController := &admin.UserScoreLogController{}
	admin_router.GET("/user-score-log/:id", adminUserScoreLogController.GetBy)
	admin_router.Any("/user-score-log/list", adminUserScoreLogController.AnyList)

	//admin_router.Party("/operate-log").Handle(new(admin.OperateLogController))
	adminOperateLogController := &admin.OperateLogController{}
	admin_router.GET("/operate-log/:id", adminOperateLogController.GetBy)
	admin_router.Any("/operate-log/list", adminOperateLogController.AnyList)

	api_router.GET("/api/img/proxy", func(ctx *gin.Context) {
		url := simple.FormValue(ctx, "url")
		resp, err := resty.New().R().Get(url)
		ctx.Header("Content-Type", "image/jpg")
		if err == nil {
			ctx.JSON(200, resp.Body())
		} else {
			logrus.Error(err)
		}
	})

}
