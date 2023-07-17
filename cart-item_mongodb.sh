#!/bin/bash

# Connect to the MongoDB server
mongosh "mongodb://localhost:27017" <<EOF

use skripsi

db.createCollection("cartitem")

db.cartitem.createIndex({ "uuid": 1 }, { unique: true })

EOF
