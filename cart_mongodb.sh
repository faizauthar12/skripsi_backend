#!/bin/bash

# Connect to the MongoDB server
mongosh "mongodb://localhost:27017" <<EOF

use skripsi

db.createCollection("cart")

db.cart.createIndex({ "uuid": 1 }, { unique: true })

db.cart.createIndex({ "customeruuid": 1 }, { unique: true })

EOF
