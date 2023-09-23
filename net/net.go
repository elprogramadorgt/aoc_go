package net


import (
	"net/http"
	"io"
)


func GetRemoteFile(url string) ([]byte, error)  {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	
	req.Header.Set("Cookie", "_ga=GA1.2.1324803828.1688699141; _ga_MHSNPJKWC7=GS1.2.1693963903.11.1.1693964221.0.0.0; _gid=GA1.2.1518597740.1693963902; session=53616c7465645f5fd323b23bf55e19d41f16fb784f22a86a1885df60cd910253dc8fabc44eb628ba7a41a57a405dfa413d4249d92d79fdbda7e38aff2b80cb0d")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}


	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}
