package handlers

import (
	"context"
	"fmt"
	"music-backend-test/internal/api/http/presenter"
	"music-backend-test/internal/entity"
	"music-backend-test/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type musicHandlers struct {
	interactor usecase.MusicInteractor
	presenter  presenter.Presenter
}

func NewMusicHandlers(interactor usecase.MusicInteractor, presenter presenter.Presenter) *musicHandlers {
	return &musicHandlers{
		interactor: interactor,
		presenter:  presenter,
	}
}

// GetAllHandler godoc
// @Summary Получение всех треков
// @Description Получение всех треков
// @Tags Music
// @Accept json
// @Produce plain
// @Security JwtAuth
// @Success 200 {object} []view.MusicView "Данные трека"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/catalog [get]
func (m *musicHandlers) GetAll(c *gin.Context) {
	ctx := context.Background()
	musics, err := m.interactor.GetAll(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.GetAll: %w", err))
		return
	}

	c.JSON(http.StatusOK, m.presenter.ToListMusicView(musics))
}

// GetAndSortByPopularHandler godoc
// @Summary Получение треков отсортированных по популярности
// @Description Получение треков отсортированных по популярности
// @Tags Music
// @Accept json
// @Produce plain
// @Security JwtAuth
// @Success 200 {object} []view.MusicView "Список треков"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/popular [get]
func (m *musicHandlers) GetAndSortByPopular(c *gin.Context) {
	ctx := context.Background()
	musics, err := m.interactor.GetAndSortByPopular(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.GetAndSortByPopular: %w", err))
	}
	c.JSON(http.StatusOK, m.presenter.ToListMusicView(musics))
}

// GetAllSortByTimeHandler godoc
// @Summary Получение треков отсортированных по популярности
// @Description Получение треков отсортированных по популярности
// @Tags Music
// @Accept json
// @Produce plain
// @Security JwtAuth
// @Success 200 {object} []view.MusicView "Список треков"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/release [get]
func (m *musicHandlers) GetAllSortByTime(c *gin.Context) {
	ctx := context.Background()
	musics, err := m.interactor.GetAllSortByTime(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.GetAllSortByTime: %w", err))
		return
	}
	c.JSON(http.StatusOK, m.presenter.ToListMusicView(musics))
}

// GetFileHandler godoc
// @Summary Скачивание файла трека
// @Description Скачивание файла трека по id трека
// @Tags Music
// @Produce json
// @Param id path string true "id трека"
// @Security JwtAuth
// @Success 200 {object} view.MusicView "Данные трека"
// @Success 200 {file} file "Файл трека"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/download/{id} [get]
func (m *musicHandlers) Get(c *gin.Context) {
	ctx := context.Background()

	musicId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't parse id: %w", err))
		return
	}

	music, err := m.interactor.Get(ctx, musicId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.Get: %w", err))
		return
	}

	c.FileAttachment(music.FilePath(), music.FileName)
}

// CreateHandler godoc
// @Summary Создание трека
// @Description Создание нового трека
// @Tags Music
// @Accept json
// @Produce plain
// @Security JwtAuth
// @Param request formData entity.MusicParse true "Данные трека"
// @Param file formData file true "Файл трека"
// @Success 201 "Трек создан"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/new [post]
func (m *musicHandlers) Create(c *gin.Context) {
	ctx := context.Background()

	var music entity.MusicParse
	err := c.Request.ParseMultipartForm(64)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't read form data: %w", err))
		return
	}

	music.Name = c.Request.FormValue("name")
	music.Release, err = time.Parse("2006-01-02", c.Request.FormValue("release"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't read parse time: %w", err))
		return
	}

	music.File, music.FileHeader, err = c.Request.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't read file: %w", err))
		return
	}
	fmt.Println("\nFILE_HEADER: ", *music.FileHeader)

	err = m.interactor.Create(ctx, &music)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.Create: %w", err))
	}
	c.JSON(http.StatusCreated, nil)
}

// UpdateHandler godoc
// @Summary Обновление трека
// @Description Обновление трека
// @Tags Music
// @Accept json
// @Produce plain
// @Security JwtAuth
// @Param id path string true "id трека"
// @Param request formData entity.MusicParse true "Данные трека"
// @Param file formData file false "Файл трека"
// @Success 200 "Трек обновлен"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/{id} [put]
func (m *musicHandlers) Update(c *gin.Context) {
	ctx := context.Background()

	var music entity.MusicParse
	var err error
	musicId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't parse id: %w", err))
		return
	}

	err = c.Request.ParseMultipartForm(64)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't read form data: %w", err))
		return
	}

	music.Name = c.Request.FormValue("name")
	music.Release, err = time.Parse("2006-01-02", c.Request.FormValue("release"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't read parse time: %w", err))
		return
	}

	music.File, music.FileHeader, err = c.Request.FormFile("file")
	if err != nil {
		if err.Error() != "http: no such file" {
			c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't read file: %w", err))
			return
		}
	}

	err = m.interactor.Update(ctx, musicId, &music)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.Update: %w", err))
	}

	c.JSON(http.StatusOK, nil)
}

// DeleteHandler godoc
// @Summary Удаление трека
// @Description Удаление трека
// @Tags Music
// @Accept json
// @Produce plain
// @Security JwtAuth
// @Param id path string true "id трека"
// @Success 204 "Трек удален"
// @Failure 400 "Некорректный запрос"
// @Failure 401 "Неавторизованный запрос"
// @Failure 404 "Пользователь не найден"
// @Failure 500 "Внутренняя ошибка сервера"
// @Router /music/{id} [delete]
func (m *musicHandlers) Delete(c *gin.Context) {
	ctx := context.Background()

	musicId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, fmt.Errorf("can't parse id: %w", err))
		return
	}

	err = m.interactor.Delete(ctx, musicId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("/usecase/music.Delete: %w", err))
	}

	c.JSON(http.StatusOK, nil)
}
