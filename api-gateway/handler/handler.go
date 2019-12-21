package handler

import (
	"encoding/json"
	"net/http"
	// api "path/to/service/proto/api"
)

func ApiCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// // call the backend service
	// apiClient := api.NewApiService("com.lcb123.srv.api", client.DefaultClient)
	// rsp, err := apiClient.Call(context.TODO(), &api.Request{
	// 	Name: request["name"].(string),
	// })
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// // we want to augment the response
	// response := map[string]interface{}{
	// 	"msg": rsp.Msg,
	// 	"ref": time.Now().UnixNano(),
	// }

	// // encode and write the response as json
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
}
