package handler

import (
	"encoding/json"
	driver "fun-blogger-backend/driver"
	library "fun-blogger-backend/library"
	models "fun-blogger-backend/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BlogHandlerGet ...
@desc handling get request of /blog
@route /blog
@method POST
@access Private
*/
func BlogHandlerGet(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)

	reqToken := r.Header.Get("x-auth-token")
	var responsesMap = map[string]string{}

	userID := driver.GetUserIDByToken(reqToken)

	query := bson.M{"userID": userID}
	blogs, _ := driver.FindBlogs(query)

	if len(blogs) == 0 {
		responsesMap["blogs"] = "Blogs is empty"
	} else {
		encodedBlogs, _ := json.Marshal(blogs)
		responsesMap["blogs"] = string(encodedBlogs)
	}

	encodedResponses, _ := json.Marshal(responsesMap)

	library.ResponseByCode(200, w, string(encodedResponses))
}

/*BlogHandlerPost ...
@desc handling post request of /blog
@route /blog
@method POST
@access Private
*/
func BlogHandlerPost(w http.ResponseWriter, r *http.Request) {
	library.SetDefaultHTTPHeader(w)
	
	var blog models.Blog
	err := blog.FromJSON(r)

	var responsesMap = map[string]string{}

	if err != nil {
		library.ResponseByCode(500, w, err.Error())
		return
	}

	if blog.Title == "" {
		responsesMap["title"] = "Title cannot be Empty"
	}

	if blog.Content == "" {
		responsesMap["content"] = "Content cannot be Empty"
	}

	if len(responsesMap) != 0 {
		responsesMap["status"] = "False"
	} else {
		reqToken := r.Header.Get("x-auth-token")
		userID := driver.GetUserIDByToken(reqToken)

		blog.ID = primitive.NewObjectID()
		blog.UserID = userID
		blog.CreatedAt = library.GetCurrentDate()
		blog.UpdatedAt = library.GetCurrentDate()

		err = driver.Insert("blogs", blog)

		if err != nil {
			library.ResponseByCode(500, w, "There's something wrong")
			return
		}

		responsesMap["status"] = "true"
	}

	encodeResponses, _ := json.Marshal(responsesMap)
	library.ResponseByCode(200, w, string(encodeResponses))
}
