package blogs_test

import (
	"testing"

	"github.com/hpp131/gblog/apps/blogs"
)

func TestCreateBlog(*testing.T){
	req := blogs.NewCreateBlogRequest()
	req.Author = "han"
	req.Content = "this is content"
	req.Title = "unit_test"
	
}

func NewCreateBlogRequest() {
	panic("unimplemented")
}

func TestQueryBlog(*testing.T){
	
}

func TestPatchBlog(*testing.T){
	
}


func TestPutBlog(*testing.T){
	
}


func TestDeleteBlog(*testing.T){
	
}