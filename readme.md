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
  - 1,000,001 - 2,000,000 อัตราภาษี 20%
  - มากกว่า 2,000,000 อัตราภาษี 35%
- เงินบริจาคสามารถหย่อนได้สูงสุด 100,000 บาท

## Non-Functional Requirement

- มี `Unit Test` ครอบคลุม
- ใช้ `go module`
- ใช้ PostgreSQL
- ใช้ go module `go mod init github.com/<your github name>/assessment-tax`
- ใช้ go 1.21 or above
- API port _MUST_ get from environment variable name `PORT` (should be able to config for API start from port `:2565`)
- database url _MUST_ get from environment variable name `DATABASE_URL`
- ใช้ `docker-compose` สำหรับต่อ Database
- API support `Graceful Shutdown`
- มี Dockerfile สำหรับ build image และเป็น `Multi-stage build`
- ใช้ `HTTP Status Code` อย่างเหมาะสม
- ใช้ `HTTP Method` อย่างเหมาะสม
- ใช้ `gofmt`
- ใช้ `go vet`
- แยก Branch ของแต่ละ Story และ Merge กลับไปยัง `main` Branch
  - เช่น `feature/story-1`

## Stories Note

- ผู้ใช้คำนวนภาษีตาม เงินได้ และฐานภาษี
- ผู้ใช้คำนวนภาษี โดยสามารถใช้ค่าลดหย่อนจากการบริจาคได้
- ผู้ใช้คำนวนภาษี โดยสามารถใช้ค่า wht เพื่อคำนวนเงินที่สามารถขอคืนได้
- แอดมินสามารถตั้งค่า ค่าลดหย่อนได้
- แสดงข้อมูลเพิ่มเติมตามขั้นบันใดภาษีได้
- ผู้ใช้สามารถคำนวนภาษีตาม csv ที่อัพโหลดมาได้

## User stories

<summary>### Story: EXP01</summary>

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
{
  "tax": 29000.0
}
```

<details>
<summary>Calculation guide</summary>

500,000 (รายรับ) - 60,0000 (ค่าลดหย่อนส่วนตัว) = 440,000

| Tax Level | Tax |
|-|-|
|0-150,000|0|
|150,001-500,000|29,000|
|500,001-1,000,000|0|
|1,000,001-2,000,000|0|
|2,000,001 ขึ้นไป|0|
</details>

<summary>### Story: EXP02</summary>

```
* As user, I want to calculate my tax with WHT
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจาก ข้อมูลที่ส่งให้ พร้อมกับข้อมูลหักภาษี ณ ที่จ่าย
```

`POST:` /calculation

```json
{
  "totalIncome": 500000.0,
  "wht": 25000.0,
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
{
  "tax": 4000.0
}
```

<details>
<summary>Calculation guide</summary>

500,000 (รายรับ) - 60,0000 (ค่าลดหย่อนส่วนตัว) = 440,000

ภาษีที่จะต้องชำระ 29,000.00 - 25,000.00 = 4,000

</details>

<!-- -- -->

<!-- <details> -->
<summary>### Story: EXP03</summary>

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
      "amount": 200000.0
    }
  ]
}
```

Response body

```json
{
  "tax": 19000.0
}
```

<details>
<summary>Calculation guide</summary>

500,000 (รายรับ) - 60,0000 (ค่าลดหย่อนส่วนตัว)  - 100,000 = 340,000

| Tax Level | Tax |
|-|-|
|0-150,000|0|
|150,001-500,000|19,000|
|500,001-1,000,000|0|
|1,000,001-2,000,000|0|
|2,000,001 ขึ้นไป|0|

</details>
<!-- </details> -->