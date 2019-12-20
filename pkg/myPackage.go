package pkg

type Files struct{
	ID string `json:"id"`
	url []string `json:"file_link"`
	start_time string `json:"start_time"`
	end_time string `json:"end_time"`
	download_type string `json:"download_type"`
	file_info []*downloadInfo `json:"file_info"`

}

type downloadInfo struct{
	file_path string `json:"file_path"`
	download_path string `json:"download_path"`
}
