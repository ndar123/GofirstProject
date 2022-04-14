package service

import "GofirstProject/entity"

type VideoService interface {
	Save() entity.Videos
	FindAll() []entity.Videos
}

type videoService struct {
	videos []entity.Videos
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Videos) entity.Videos {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Videos {
	return service.videos

}
