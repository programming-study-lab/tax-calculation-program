# Tax Calculator API

REST API สำหรับคำนวณภาษีเงินได้บุคคลธรรมดา

## เทคโนโลยีที่ใช้
- Golang 1.26.3
- Gin v1.12.0 
- Postman

## วิธีการติดตั้ง
```bash
go get github.com/gin-gonic/gin
```

## วิธีการรัน
```bash
go run main.go
```
API จะรันที่ `http://localhost:5000`

## API Endpoints

### POST /tax/calculations
คำนวณภาษีเงินได้

**Request Body:**
```json
{
    "totalIncome": 750000,
    "wht": 0,
    "allowances": []
}
```

**Response:***
```json
{
    "tax": 63500,
    "taxLevel": [
        {
            "level": "0-150,000",
            "tax": 0
        },
        {
            "level": "150,001-500,000",
            "tax": 35000
        },
        {
            "level": "500,001-1,000,000",
            "tax": 28500
        },
        {
            "level": "1,000,001-2,000,000",
            "tax": 0
        },
        {
            "level": "มากกว่า 2,000,000",
            "tax": 0
        }
    ]
}
```

## ตัวอย่างการใช้งาน

### คำนวณภาษีพื้นฐาน
```bash
curl -X POST http://localhost:5000/tax/calculations \
  -H "Content-Type: application/json" \
  -d '{
    "totalIncome": 750000,
    "wht": 0,
    "allowances": []
  }'
```

### คำนวณภาษีพร้อม WHT
```bash
curl -X POST http://localhost:5000/tax/calculations \
  -H "Content-Type: application/json" \
  -d '{
    "totalIncome": 600000,
    "wht": 15000,
    "allowances": []
  }'
```