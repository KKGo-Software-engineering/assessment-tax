# K-Tax โปรแกรมคำนวนภาษี

K-Tax เป็น Application คำนวนภาษี ที่ให้ผู้ใช้งานสามารถคำนวนภาษีบุคคลธรรมดา ตามขั้นบันใดภาษี พร้อมกับคำนวนค่าลดหย่อน และภาษีที่ต้องได้รับคืน

## Getting Started

1. Clone this repository
2. Create repository in your `Github` account
3. Push starter code to your `Github` account
4. Create new branch that according to story name
5. Follow `Functional Requirement` and `Non-Functional Requirement`

## Functional Requirement

- ผู้ใช้งาน สามารถส่งข้อมูลเพื่อคำนวนภาษีได้
- ผู้ใช้งาน แสดงภาษีที่ต้องจ่ายหรือได้รับในปีนั้น ๆ ได้
- การคำนวนภาษีคำนวนจาก เงินหัก ณ ที่จ่าย / ค่าลดหย่อนส่วนตัว/ขั้นบันใดภาษี/เงินบริจาค
- การคำนวนภาษีตามขั้นบันใด
  - รายได้ 0 - 150,000 ได้รับการยกเว้น
  - 150,001 - 500,000 อัตราภาษี 10%
  - 500,001 - 1,000,000 อัตราภาษี 15%
  - 1,000,001 - 2,000,000 อัตราภาษี 20%
  - มากกว่า 2,000,000 อัตราภาษี 35%
- เงินบริจาคสามารถหย่อนได้สูงสุด 100,000 บาท
- ค่าลดหย่อนส่วนตัวมีค่าเริ่มต้นที่ 60,000 บาท
- k-receipt โครงการช้อปลดภาษี ซึ่งสามารถลดหย่อนได้สูงสุด 50,000 บาทเป็นค่าเริ่มต้น
- แอดมิน สามารถกำหนดค่าลดหย่อนส่วนตัวได้โดยไม่เกิน 100,000 บาท
- แอดมิน สามารถกำหนด k-receipt สูงสุดได้ แต่ไม่เกิน 100,000 บาท
- ค่าลดหย่อนส่วนตัวต้องมีค่ามากกว่า 10,000 บาท
- ค่าลด k-receipt ต้องมีค่ามากกว่า 0 บาท
- ในกรณีที่รายรับ รวมหักค่าลดหย่อน พร้อมทั้ง wht พบว่าต้องได้เงินคืน จะต้องคำนวนเงินที่ต้องได้รับคืนใน field ใหม่ ที่ชื่อว่า taxRefund

## Non-Functional Requirement
- มี `Unit Test` ครอบคลุม
- ใช้ `go module`
- ใช้ go module `go mod init github.com/<your github name>/assessment-tax`
- ใช้ go 1.21 or above
- ใช้ `PostgreSQL`
- API port _MUST_ get from `environment variable` name `PORT`
- database url _MUST_ get from environment variable name `DATABASE_URL`
  - ตัวอย่าง `DATABASE_URL=host={REPLACE_ME} port=5432 user={REPLACE_ME} password={REPLACE_ME} dbname={REPLACE_ME} sslmode=disable`
- ใช้ `docker compose` สำหรับต่อ Database
- API support `Graceful Shutdown`
  - เช่น ถ้ามีการกด `Ctrl + C` จะ print `shutting down the server`
- มี Dockerfile สำหรับ build image และเป็น `Multi-stage build`
- ใช้ `HTTP Method` และ `HTTP Status Code` อย่างเหมาะสม
- ใช้ `gofmt` และ `go vet`
- แยก Branch ของแต่ละ Story ออกจาก `main` และ Merge กลับไปยัง `main` Branch เสมอ
  - เช่น story ที่ 1 จะใช้ branch ชื่อ `feature/story-1` หรือ `feature/store-1-create-tax-calculation`
- admin กำหนด Basic authen ด้วย username: `adminTax`, password: `admin!`
  - username และ password ต้องเป็น environment variable
  - และ `env` ต้องเป็นชื่อ `ADMIN_USERNAME` และ `ADMIN_PASSWORD`
- **การ run program จะใช้คำสั่ง docker compose up เพื่อเตรียม environment และ go run main.go เพื่อ start api**
  - **หากต้องมีการใช้คำสั่งอื่น ๆ เพื่อทำให้โปรแกรมทำงานได้ จะไม่นับคะแนนหรือถูกหักคะแนน**
  - การตรวจจะทำการ export `env` ไว้ล่วงหน้าก่อนรัน ดังนี้
	- `export PORT=8080`
	- `export DATABASE_URL={REPLACE_ME}`
	- `export ADMIN_USERNAME=adminTax`
	- `export ADMIN_PASSWORD=admin!`
- port ของ api จะต้องเป็น 8080

## Assumption

- รองรับแค่ปีเดียวคือ 2567
- ไม่มีเก็บข้อมูลภาษีของผู้ใช้งาน
- อัตราภาษีไม่มีการเปลี่ยนแปลงในอนาคต
- ค่าลดหย่อนมีได้ 3 ชนิดเท่านั้น ค่าลดหย่อนส่วนตัว/เงินบริจาค/ช้อปปลดภาษี
- ค่าลดหย่อนที่จะส่งเข้ามาคำนวนไม่มีค่าน้อยกว่า 0
- ข้อมูล wht ที่จะถูกส่งเข้ามาคำนวน ไม่สามารถมีค่าน้อยกว่า 0 หรือมากกว่ารายรับได้
- csv ที่รับเข้ามา ต้องใช้ชื่อตามที่กำหนดให้ และมีโครงสร้างข้อมูลตามตัวอย่างเท่านั้น
- ข้อมูลที่รับเข้ามา ต้องผ่านการตรวจสอบความถูกต้องและความสมบูรณ์ก่อนการคำนวน

## Stories Note

- ผู้ใช้คำนวนภาษีตาม เงินได้ และฐานภาษี
- ผู้ใช้คำนวนภาษี โดยสามารถใช้ค่าลดหย่อนจากการบริจาคได้
- ผู้ใช้คำนวนภาษี โดยสามารถใช้ค่า wht เพื่อคำนวนเงินที่สามารถขอคืนได้
  - (wht: with holding tax หมายถึงเงินจำนวนนึงที่ต้องหักไว้ ณ ที่จ่ายเช่น รายรับ 1 ครั้งมี wht 5% แปลว่าได้รับเงิน 10,000 บาทจะต้องถูกหัก 500 บาท แล้วเงินส่วนนี้จะถูกส่งเข้าระบบ เสมือนได้ชำระภาษีล่วงหน้าแล้ว ถ้ารายได้ไม่ถึงเกณฑ์ที่ต้องเสียเพิ่มเติม สามารถขอคืนได้)
- แอดมินสามารถตั้งค่า ค่าลดหย่อนได้
- แสดงข้อมูลเพิ่มเติมตามขั้นบันใดภาษีได้
- ผู้ใช้สามารถคำนวนภาษีตาม CSV ที่อัพโหลดมาได้

## User stories

### Story: EXP01

```
* As user, I want to calculate my tax
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจากข้อมูลที่ส่งให้
```

`POST:` tax/calculations

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

-------
### Story: EXP02

```
* As user, I want to calculate my tax with WHT
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจาก ข้อมูลที่ส่งให้ พร้อมกับข้อมูลหักภาษี ณ ที่จ่าย
```

`POST:` tax/calculations

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

-------
### Story: EXP03

```
* As user, I want to calculate my tax
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจากข้อมูลที่ส่งให้และข้อมูลค่าลดหย่อน
```

`POST:` tax/calculations

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

500,000 (รายรับ) - 60,0000 (ค่าลดหย่อนส่วนตัว) - 100,000 (เงินบริจาค) = 340,000

| Tax Level | Tax |
|-|-|
|0-150,000|0|
|150,001-500,000|19,000|
|500,001-1,000,000|0|
|1,000,001-2,000,000|0|
|2,000,001 ขึ้นไป|0|
----
</details>


-------
### Story: EXP04

```
* As user, I want to calculate my tax with tax level detail
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจาก ข้อมูลที่ส่งให้ พร้อมระบุรายละเอียดของขั้นบันใดภาษี
```

`POST:` tax/calculations

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
  "tax": 19000.0,
  "taxLevel": [
    {
      "level": "0-150,000",
      "tax": 0.0
    },
    {
      "level": "150,001-500,000",
      "tax": 19000.0
    },
    {
      "level": "500,001-1,000,000",
      "tax": 0.0
    },
    {
      "level": "1,000,001-2,000,000",
      "tax": 0.0
    },
    {
      "level": "2,000,001 ขึ้นไป",
      "tax": 0.0
    }
  ]
}
```
----

### Story: EXP05

```
* As admin, I want to setting personal deduction
ในฐานะ Admin ฉันต้องการตั้งค่าลดหย่อนส่วนตัว
```

`POST:` /admin/deductions/personal

```json
{
  "amount": 70000.0
}
```

Response body

```json
{
  "personalDeduction": 70000.0
}
```
----


### Story: EXP06

```
* As user, I want to calculate my tax with csv
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีด้วยข้อมูลที่ upload เป็น csv และมี validation เช่น ใส่ empty ไม่ได้ หรือ ใส่ข้อมูลผิด format ไม่ได้
```

`POST:` tax/calculations/upload-csv

form-data:
  - taxFile: taxes.csv

```
totalIncome,wht,donation
500000,0,0
600000,40000,20000
750000,50000,15000
```

Response body

```json
{
  "taxes": [
    {
      "totalIncome": 500000.0,
      "tax": 29000.0
    },
    ...
  ]
}
```

-------
### Story: EXP07

```
* As user, I want to calculate my tax with tax level detail
ในฐานะผู้ใช้ ฉันต้องการคำนวนภาษีจาก ข้อมูลที่ส่งให้พร้อมค่าลดหย่อน พร้อมระบุรายละเอียดของขั้นบันใดภาษี
```

`POST:` tax/calculations

```json
{
  "totalIncome": 500000.0,
  "wht": 0.0,
  "allowances": [
    {
      "allowanceType": "k-receipt",
      "amount": 200000.0
    },
    {
      "allowanceType": "donation",
      "amount": 100000.0
    }
  ]
}
```

Response body

```json
{
  "tax": 14000.0,
  "taxLevel": [
    {
      "level": "0-150,000",
      "tax": 0.0
    },
    {
      "level": "150,001-500,000",
      "tax": 14000.0
    },
    {
      "level": "500,001-1,000,000",
      "tax": 0.0
    },
    {
      "level": "1,000,001-2,000,000",
      "tax": 0.0
    },
    {
      "level": "2,000,001 ขึ้นไป",
      "tax": 0.0
    }
  ]
}
```
<details>
<summary>Calculation guide</summary>

500,000 (รายรับ) - 60,0000 (ค่าลดหย่อนส่วนตัว) - 100,000 (เงินบริจาค) - 50,000 (k-receipt) = 290,000

| Tax Level | Tax    |
|-|--------|
|0-150,000| 0      |
|150,001-500,000| 14,000 |
|500,001-1,000,000| 0      |
|1,000,001-2,000,000| 0      |
|2,000,001 ขึ้นไป| 0      |
----
</details>

----

### Story: EXP08

```
* As admin, I want to setting k-receipt deduction
ในฐานะ Admin ฉันต้องการตั้งค่า k-receipt สำหรับลดหย่อน
```

`POST:` /admin/deductions/k-receipt

```json
{
  "amount": 70000.0
}
```

Response body

```json
{
  "kReceipt": 70000.0
}
```
----
