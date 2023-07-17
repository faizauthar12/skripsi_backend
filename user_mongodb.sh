#!/bin/bash

# Connect to the MongoDB server
mongosh "mongodb://localhost:27017" <<EOF

use skripsi

db.createCollection("user")

db.user.createIndex({ "uuid": 1 }, { unique: true })

db.user.createIndex({ "email": 1 }, { unique: true })

EOF
