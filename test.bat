###
POST http://localhost:8080/sotrudnik/department
Accept: application/json

{
  "departmentName": "Sales"
}

###

POST http://localhost:8080/sotrudnik/company
Accept: application/json

{
  "companyId": 41
}

###
DELETE http://localhost:8080/sotrudnik
Accept: application/json

{
  "id": 2
}

###
PUT http://localhost:8080/sotrudnik
Accept: application/json

{
  "name": "Olya",
  "surname": "Volgina",
  "phone": "+71234567890",
  "companyId": 41,
  "passport": {
    "type": "International",
    "number": "AB1234567"
  },
  "department": {
    "name": "Sales",
    "phone": "+70987654321"
  }
}


###
