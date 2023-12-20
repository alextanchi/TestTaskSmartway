###
POST http://localhost:8080/sotrudnik/department
Accept: application/json

{
  "departmentName": "Sales"
}

<> 2023-12-20T215152.200.txt

###

POST http://localhost:8080/sotrudnik/company
Accept: application/json

{
  "companyId": 456
}

<> 2023-12-20T210407.200.txt

###
DELETE http://localhost:8080/sotrudnik
Accept: application/json

{
  "id": 2
}

<> 2023-12-20T210611.404.txt
<> 2023-12-20T210532.404.txt
<> 2023-12-20T210520.404.txt
<> 2023-12-20T203629.404.txt
<> 2023-12-20T203453.500.txt
<> 2023-12-20T203245.500.txt
<> 2023-12-20T203051.500.txt
###
PUT http://localhost:8080/sotrudnik
Accept: application/json

{
  "name": "Olya",
  "surname": "Volgina",
  "phone": "+71234567890",
  "companyId": 456,
  "passport": {
    "type": "International",
    "number": "AB1234567"
  },
  "department": {
    "name": "Sales",
    "phone": "+70987654321"
  }
}

<> 2023-12-20T215133.200.txt
<> 2023-12-20T215048.200.txt
<> 2023-12-20T210225.200.txt
<> 2023-12-20T200749.200.txt
<> 2023-12-20T200549.200.txt
###
