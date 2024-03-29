package repository

import (
	"context"
	"fmt"
	"music-backend-test/internal/db"
	"music-backend-test/internal/entity"
	"music-backend-test/internal/utils"

	"github.com/google/uuid"
)

type musicRepository struct {
	source     db.MusicSource
	utils      utils.MusicUtils
	FileSystem utils.FileSystem
}

func NewMusicRepository(source db.MusicSource, utils utils.MusicUtils, filesystem utils.FileSystem) *musicRepository {
	return &musicRepository{
		source:     source,
		utils:      utils,
		FileSystem: filesystem,
	}
}

func (m *musicRepository) GetAll(ctx context.Context) ([]*entity.MusicDB, error) {
	musicsDB, err := m.source.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("/db/music.GetAll: %w", err)
	}
	return musicsDB, nil
}

func (m *musicRepository) Get(ctx context.Context, musicId uuid.UUID) (*entity.MusicDB, error) {
	musicDB, err := m.source.Get(ctx, musicId)
	if err != nil {
		return nil, fmt.Errorf("/db/music.Get: %w", err)
	}

	return musicDB, nil
}

func (m *musicRepository) GetAndSortByPopular(ctx context.Context) ([]*entity.MusicDB, error) {
	musicsDB, err := m.source.GetAndSortByPopular(ctx)
	if err != nil {
		return nil, fmt.Errorf("/db/music.GetAndSortByPopular: %w", err)
	}

	return musicsDB, nil
}

func (m *musicRepository) GetAllSortByTime(ctx context.Context) ([]*entity.MusicDB, error) {
	musicsDB, err := m.source.GetAllSortByTime(ctx)
	if err != nil {
		return nil, fmt.Errorf("/db/music.GetAllSortByTime: %w", err)
	}

	return musicsDB, nil
}

func (m *musicRepository) Create(ctx context.Context, musicParse *entity.MusicParse) error {
	musicCreate := &entity.MusicDB{
		Name:     musicParse.Name,
		Release:  musicParse.Release,
		FileName: musicParse.FileHeader.Filename,
	}

	fileType, err := m.utils.GetSupportedFileType(musicParse.FileHeader.Filename)
	if err != nil {
		return fmt.Errorf("/utils.GetSupportedFileType: %w", err)
	}

	// скачиваем файл
	download_file, err := m.FileSystem.Create(musicCreate.FilePath())
	if err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}
	defer download_file.Close()
	defer musicParse.File.Close()

	if _, err := m.FileSystem.Copy(download_file, musicParse.File); err != nil {
		return fmt.Errorf("can't copy file: %w", err)
	}

	musicCreate.Size = uint64(musicParse.FileHeader.Size)

	musicCreate.Duration, err = m.utils.GetAudioDuration(fileType, musicCreate.FilePath(), m.FileSystem)
	if err != nil {
		return fmt.Errorf("/utils.GetAudioDuration: %w", err)
	}

	err = m.source.Create(ctx, musicCreate)
	if err != nil {
		return fmt.Errorf("/db/music.Create: %w", err)
	}

	return nil
}

func (m *musicRepository) Update(ctx context.Context, id uuid.UUID, musicParse *entity.MusicParse) error {
	musicUpdate := &entity.MusicDB{
		Id:      id,
		Name:    musicParse.Name,
		Release: musicParse.Release,
	}

	if musicParse.FileHeader != nil {
		musicUpdate.FileName = musicParse.FileHeader.Filename

		fileType, err := m.utils.GetSupportedFileType(musicParse.FileHeader.Filename)
		if err != nil {
			return fmt.Errorf("/utils.GetSupportedFileType: %w", err)
		}

		music, err := m.source.Get(ctx, id)
		if err != nil {
			return fmt.Errorf("/db/music.Get: %w", err)
		}

		m.FileSystem.Remove(music.FilePath())

		download_file, err := m.FileSystem.Create(music.FilePath())
		if err != nil {
			return fmt.Errorf("can't create file: %w", err)
		}
		defer download_file.Close()
		defer musicParse.File.Close()

		if _, err := m.FileSystem.Copy(download_file, musicParse.File); err != nil {
			return fmt.Errorf("can't copy file: %w", err)
		}

		musicUpdate.Size = uint64(musicParse.FileHeader.Size)

		musicUpdate.Duration, err = m.utils.GetAudioDuration(fileType, musicUpdate.FilePath(), m.FileSystem)
		if err != nil {
			return fmt.Errorf("/utils.GetAudioDuration: %w", err)
		}

		musicUpdate.FileName = musicParse.FileHeader.Filename
	}

	err := m.source.Update(ctx, musicUpdate)
	if err != nil {
		return fmt.Errorf("/db/music.Update: %w", err)
	}

	return nil
}

func (m *musicRepository) Delete(ctx context.Context, id uuid.UUID) error {
	music, err := m.source.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("/db/music.Get: %w", err)
	}
	err = m.FileSystem.Remove(music.FilePath())
	if err != nil {
		return fmt.Errorf("can't delete music file: %w", err)
	}

	err = m.source.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("/db/music.Delete: %w", err)
	}

	return nil
}
