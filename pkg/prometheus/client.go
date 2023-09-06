package prometheus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"zmos-underground-monitor/conf"
)

var (
	prometheusAddr string
	query          string
)

type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type PrometheusRangeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Values [][]interface{}   `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

func queryPrometheus(url string, response interface{}) error {
	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse JSON response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return nil
}

func GetPoint(query string) (value float64, err error) {
	//单个指标查询
	queryUrl := fmt.Sprintf("http://%s/api/v1/query?query=%s", conf.Conf.Prometheus.URL, query)
	// Query and parse PrometheusResponse
	var promResp PrometheusResponse
	err = queryPrometheus(queryUrl, &promResp)
	if err != nil {
		fmt.Println("Error querying Prometheus:", err)
		return
	}

	// Print the result
	if promResp.Status == "success" {
		for _, result := range promResp.Data.Result {
			metric := result.Metric
			if len(result.Value) == 2 {
				//value[0] 是当前时间
				//value[1] 是当前值
				value, err = strconv.ParseFloat(result.Value[1].(string), 64)
				if err != nil {
					fmt.Printf("类型转换错误：%v", result.Value[1])
					return
				}
			}
			fmt.Printf("Metric: %v\nValue: %v\n", metric, value)
		}
	} else {
		fmt.Println("Prometheus query failed:", promResp.Status)
		err = errors.New(fmt.Sprintf("获取指标%s错误", query))
	}
	return value, err
}
