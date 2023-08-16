package net


import (
	"net/http"
	"io"
)


func GetRemoteFile(url string) ([]byte, error)  {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	
	req.Header.Set("Cookie", "_ga=GA1.2.1324803828.1688699141; _ga_MHSNPJKWC7=GS1.2.1690335385.2.1.1690336637.0.0.0; _gid=GA1.2.2109103212.1690335385; session=53616c7465645f5fa358497347a4c745e807e0e5240d936cb580425af79ec84c6798175e9a2356fd29b112f6345183c86a11e5cd81f0d1a5ecf774c1af89b663")
	

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
