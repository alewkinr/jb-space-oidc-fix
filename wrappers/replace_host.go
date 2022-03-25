package wrappers

import (
	"net/http"
)

// MakeRequest — собираем корректный запрос
func MakeRequest(targetHost string, req *http.Request) (*http.Request, error) {
	outreq, makeRequestErr := http.NewRequest("GET", targetHost, nil)
	if makeRequestErr != nil {
		return nil, makeRequestErr
	}

	setTokenInHeaderErr := SetBearerTokenFromQuery(targetHost, outreq)
	if setTokenInHeaderErr != nil {
		return nil, setTokenInHeaderErr
	}

	return outreq, nil
	//req.Header.Add("cookie", "AMAuthCookie=AQIC5wM2LY4SfczOSN26FAwFRT9D2_GiZQbBoQR_QnIDDMg.*AAJTSQACMDQAAlNLABMyMjk0MjY5NzM3Mzk0NzU4OTA0AAJTMQACMTM.*; amlbcookie=13")
	//req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI2THdqZmo5UkdhallmNnRuTl9SeVVoMjdRajE2Vm5uVzRwRmVUQVRIekJFIn0.eyJleHAiOjE2NDgxOTczMDMsImlhdCI6MTY0ODE5NjQwMywiYXV0aF90aW1lIjoxNjQ4MTk2MzIwLCJqdGkiOiIwMzBhYzZmMi1jNjc0LTQ0MjQtOTM5Yi1jYmYxNDI0OGYyNDEiLCJpc3MiOiJodHRwczovL2lzc28ubXRzLnJ1L2F1dGgvcmVhbG1zL210cyIsImF1ZCI6WyJ2YXVsdC1kZXYiLCJhY2NvdW50Il0sInN1YiI6ImQ3OTRjNTAyLWMzODMtNDZkZS1hNGQ5LTEzNDYwMTdiYmM3NiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImpiLXNwYWNlIiwic2Vzc2lvbl9zdGF0ZSI6IjU0OTNjMmY3LWJhMzYtNGZjMy1hNWM3LTZiYzU3Y2IyYTdmZSIsImFjciI6IjAiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9zcGFjZS5pbmZyYS5jbG91ZC5tdHMucnUiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ2YXVsdC1kZXYiOnsicm9sZXMiOlsiZGVmYXVsdCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIGZ1bGxuYW1lIHVzZXJQcmluY2lwYWxOYW1lIGVtcGxveWVlSUQtaWQgcHJvZmlsZS1pZCB1c2VybmFtZSBlbWFpbC1pZCBlbXBsb3llZUlEIHByb2ZpbGUgdXNlclByaW5jaXBhbE5hbWUtaWQgZW1haWwgdXNlcm5hbWUtaWQgZnVsbG5hbWUtaWQiLCJzaWQiOiI1NDkzYzJmNy1iYTM2LTRmYzMtYTVjNy02YmM1N2NiMmE3ZmUiLCJuYW1lIjoi0JTQsNC90LjQuyDQodC10LvQuNGE0LDQvdC-0LIiLCJlbXBsb3llZUlEIjoiNzI0MDQzIiwiZnVsbG5hbWUiOiLQodC10LvQuNGE0LDQvdC-0LIg0JTQsNC90LjQuyDQk9C10L7RgNCz0LjQtdCy0LjRhyIsInByZWZlcnJlZF91c2VybmFtZSI6ImRnc2VsaWYxIiwibWlkZGxlX25hbWUiOiLQk9C10L7RgNCz0LjQtdCy0LjRhyIsImdpdmVuX25hbWUiOiLQlNCw0L3QuNC7IiwiZmFtaWx5X25hbWUiOiLQodC10LvQuNGE0LDQvdC-0LIiLCJlbWFpbCI6ImRnc2VsaWYxQG10cy5ydSIsInVzZXJQcmluY2lwYWxOYW1lIjoiZGdzZWxpZjFAbXRzLnJ1IiwidXNlcm5hbWUiOiJkZ3NlbGlmMSJ9.aeCLLhuqrtVgVLDLEv5I2bMcPCs_SHGlyheY5PzZC4TlK5wuuUtPAwvW6L20e9qVyvCxT-kE72lW88bJ1_rDSwighUwXVIOifIHQDzu55wKc3tN6S7E1HszB7_ovjiJHy7GxgPN0ZNxH8PkWnL8NvWQj7qpC2xLD2t2wSfuvSuSa_0OIp-sFdQBMgGsUaV8djZ4mP0KEFMaGtgUR8noDJ8N-5wYjbhwZdB_b4-ACOJ7Ez2L7umryAtpFCnXL951P2MZKOvGu3xkfLiConvs8nccoVrOij8Qna_zXuTLaw3M4CBRT8br9o82kBK8SaO_OQGh2DH2VzTT8eLW4cGWwEA")
}
