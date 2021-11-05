package delivery

import (
	"CleanTodo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
	TodoUsecase model.TodoUsecase
}

func NewTodoRouter(g *gin.Engine, t model.TodoUsecase) {
	router := &TodoRouter{
		TodoUsecase: t,
	}
	g.POST("/api/todo", router.TodoCreate)
	g.DELETE("/api/todo/:listid", router.TodoDelete)
	g.PUT("/api/todo/:listid", router.TodoUpdate)
	g.GET("/api/todo/", router.TodoGetAll)
	g.GET("/api/todo/:listid", router.TodoGetOne)

	g.Run(":8887")
}

func (t *TodoRouter) TodoCreate(g *gin.Context) {
	info := model.List{}
	g.BindJSON(&info)
	title := info.Title
	content := info.Content
	listid, err := t.TodoUsecase.Create(title, content)
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"listid":  listid,
			"message": "create fail!",
			"error":   err.Error(),
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"listid":  listid,
			"title":   title,
			"content": content,
			"message": "create success!",
		})
	}
}

func (t *TodoRouter) TodoDelete(g *gin.Context) {
	listid := g.Param("listid")
	err := t.TodoUsecase.Delete(listid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "listid: " + listid + " delete fail!",
			"error":   err.Error(),
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"message": "listid: " + listid + " delete success! ",
		})
	}
}

func (t *TodoRouter) TodoUpdate(g *gin.Context) {
	listid := g.Param("listid")
	info := model.List{}
	g.BindJSON(&info)
	title := info.Title
	content := info.Content
	err := t.TodoUsecase.Update(listid, title, content)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "listid:" + listid + " update fail!",
			"error":   err.Error(),
		})
	} else {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "listid:" + listid + " update success!",
		})
	}
}

func (t *TodoRouter) TodoGetAll(g *gin.Context) {
	result, err := t.TodoUsecase.GetAll()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all list fail!",
			"error":   err.Error(),
		})
	} else {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all list success!",
			"output":  result,
		})
	}
}

func (t *TodoRouter) TodoGetOne(g *gin.Context) {
	listid := g.Param("listid")
	result, err := t.TodoUsecase.GetOne(listid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "get listid: " + listid + " fail!",
			"error":   err.Error(),
		})
	} else {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "get listid: " + listid + " success!",
			"output":  result,
		})
	}
}
