package service

import "forum/internal/model"

type Post interface {
	GetAllPosts() ([]model.PostRepresentation, error)
	CreatePost(heading string, text string, userId uint) (uint, error)
	GetPostById(postId uint) (*model.PostRepresentation, error)
	AddCategoryToPost(categoryId uint, postId uint) (uint, error)
	FilterAllPosts(filterBy string) ([]model.PostRepresentation, error)
	PersonalFilter(filterBy string, userId uint) ([]model.PostRepresentation, error)
}

type PostService struct {
	repo Post
}

func NewPostService(repo Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (ps *PostService) GetAllPosts() ([]model.PostRepresentation, error) {
	return ps.repo.GetAllPosts()
}

func (ps *PostService) CreatePost(heading string, text string, userId uint) (uint, error) {
	return ps.repo.CreatePost(heading, text, userId)
}

func (ps *PostService) GetPostById(postId uint) (*model.PostRepresentation, error) {
	return ps.repo.GetPostById(postId)
}

func (ps *PostService) AddCategoryToPost(categoryId uint, postId uint) (uint, error) {
	return ps.repo.AddCategoryToPost(categoryId, postId)
}

func (ps *PostService) FilterAllPosts(filterBy string) ([]model.PostRepresentation, error) {
	return ps.repo.FilterAllPosts(filterBy)
}

func (ps *PostService) PersonalFilter(filterBy string, userId uint) ([]model.PostRepresentation, error) {
	return ps.repo.PersonalFilter(filterBy, userId)
}
