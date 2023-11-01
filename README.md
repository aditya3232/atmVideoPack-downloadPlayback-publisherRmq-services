
# API Spec Publisher AtmVideoPack Status Mc Detection

## 1.1 Status Mc Detection

### 1.1.1 POST :: Status Mc Detection

Request :
- Method : POST
- Endpoint : `localhost:3535/publisher/atmvideopack/v1/statusmcdetection/create`
- Header :
    - Content-Type : application/x-www-form-urlencoded
    - x-api-key : required
- Body (form-data: x-www-form-urlencoded) :
    - tid : string, required
    - date_time : string, required
    - status_signal : string, required
    - status_storage : string, required
    - status_ram : string, required
    - status_cpu : string, required

- Response :

```json 
{
    "meta": {
        "message": "The request was processed successfully.",
        "code": 200
    },
    "data": {
        "tid": "16007",
        "date_time": "2023-10-09 05:12:39",
        "status_signal": "good",
        "status_storage": "good",
        "status_ram": "good",
        "status_cpu": "good"
    }
}
```






