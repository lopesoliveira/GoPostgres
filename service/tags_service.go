package service

import (
	/*"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"*/
	"github.com/lopesoliveira/GoPostgres/data/request"
	"github.com/lopesoliveira/GoPostgres/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
