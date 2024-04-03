# K-Tax โปรแกรมคำนวนภาษี

K-Tax เป็น Application คำนวนภาษี ที่ให้ผู้ใช้งานสามารถคำนวนภาษีบุคคลธรรมดา ตามขั้นบันใดภาษี พร้อมกับคำนวนค่าลดหย่อน และภาษีที่ต้องได้รับคืน

## Functional Requirement

- ผู้ใช้งาน สามารถส่งข้อมูลเพื่อคำนวนภาษีได้ (รองรับแค่ปี 2567)
- ผู้ใช้งาน แสดงภาษีที่ต้องจ่ายหรือได้รับในปีนั้น ๆ ได้
- การคำนวนภาษีคำนวนจาก เงินหัก ณ ที่จ่าย / ค่าลดหย่อนส่วนตัว/ขั้นบันใดภาษี/เงินบริจาค
- แอดมิน สามารถกำหนดค่าลดหย่อนส่วนตัวได้ โดยค่าเริ่มต้นที่ 60,000 บาท
- การคำนวนภาษีตามขั้นบันใด
  - รายได้ 0 - 150,000 ได้รับการยกเว้น
  - 150,001 - 500,000 อัตราภาษี 10%
  - 500,001 - 1,000,000 อัตราภาษี 15%
  - 1,000,001 - 1,000,000 อัตราภาษี 20%
  - มากกว่า 1,000,000 อัตราภาษี 35%
- เงินบริจาคสามารถหย่อนได้สูงสุด 100,000 บาท

## Non-Functional Requirement

- `Unit Test` is a must
- ใช้ `go module`
- ใช้ PostgreSQL
- ใช้ go module `go mod init github.com/<your github name>/assessment-tax`
- ใช้ go 1.21 or above
- api port *MUST* get from environment variable name `PORT` (should be able to config for api start from port `:2565`)
- database url *MUST* get from environment variable name `DATABASE_URL`
- Uใช้se `docker-compose` สำหรับต่อ Database
- API support `Graceful Shutdown`
- มี Dockerfile สำหรับ build image และเป็น `Multi-stage build`
- ใช้ `HTTP Status Code` อย่างเหมาะสม
- ใช้ `gofmt`
- ใช้ `go vet`

## Stories Note

- ผู้ใช้คำนวนภาษีตาม เงินได้ และฐานภาษี
- ผู้ใช้คำนวนภาษี โดยสามารถใช้ค่าลดหย่อนจากการบริจาคได้
- ผู้ใช้คำนวนภาษี โดยสามารถใช้ค่า wht เพื่อคำนวนเงินที่สามารถขอคืนได้
- แอดมินสามารถตั้งค่า ค่าลดหย่อนได้
- แสดงข้อมูลเพิ่มเติมตามขั้นบันใดภาษีได้
- ผู้ใช้สามารถคำนวนภาษีตาม csv ที่อัพโหลดมาได้

## User stories

### Story 1: EXP01

```
* As user, I want to calculate my tax
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจาก ข้อมูลที่ส่งให้
```

`POST:` /calculation

```json
{
  "totalIncome": 500000.0,
  "wht": 0.0,
  "allowances": [
    {
      "allowanceType": "donation",
      "amount": 0.0
    }
  ]
}
```

Response body

```json

```
