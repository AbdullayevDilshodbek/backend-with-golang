package db

import (
	"context"
	"testing"

	"github.com/AbdullayevDilshodbek/backend-with-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	arg := CreatePostParams{
		Title:  util.RandomString(6),
		Body:   util.RandomString(30),
		UserID: CreateRandomUser(t).ID,
		Status: "t",
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Body, post.Body)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Status, post.Status)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func createRandomPosts(t *testing.T, user User) [5]Post {
	var posts [5]Post
	for i := 0; i < 5; i++ {
		arg := CreatePostParams{
			Title:  util.RandomString(6),
			Body:   util.RandomString(30),
			UserID: user.ID,
			Status: "t",
		}

		post, err := testQueries.CreatePost(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, post)
		require.Equal(t, arg.Title, post.Title)
		require.Equal(t, arg.Body, post.Body)
		require.Equal(t, arg.UserID, post.UserID)
		require.Equal(t, arg.Status, post.Status)

		require.NotZero(t, post.ID)
		require.NotZero(t, post.CreatedAt)
		posts[i] = post
	}
	return posts
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetListPosts(t *testing.T) {
	var user User = CreateRandomUser(t)
	markedPosts := createRandomPosts(t, user)

	arg := GetPostsListParams{
		Limit:  5,
		Offset: 0,
	}

	posts, err := testQueries.GetPostsList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	for index, post := range posts {
		require.NotEmpty(t, post)
		require.Equal(t, markedPosts[4-index].Title, post.Title)
		require.Equal(t, markedPosts[4-index].Body, post.Body)
		require.Equal(t, user.ID, post.UserID)
		require.Equal(t, markedPosts[4-index].Status, post.Status)
	}
}
