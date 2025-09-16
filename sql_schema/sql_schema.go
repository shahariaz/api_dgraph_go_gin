package sqlschema

func keepData() {
	SqlSchema := `# -------------------
# Customer
# -------------------
type Customer {
  Customer.date_of_birth: date
  Customer.email: string
  Customer.gender: string
  Customer.last_login_at: datetime
  Customer.phone: string
  Customer.created_at: datetime
  Customer.deleted_at: datetime
  Customer.watch_histories: [WatchHistory]
  Customer.purchases: [Purchase]
  Customer.device_tokens: [DeviceToken]
  Customer.device_authorizations: [DeviceAuthorization]
  Customer.payments: [Payment]
}

# -------------------
# Watch History
# -------------------
type WatchHistory {
  WatchHistory.customer: Customer
  WatchHistory.content_id: int
  WatchHistory.content_type: string
  WatchHistory.deleted_at: datetime
}

# -------------------
# Purchase
# -------------------
type Purchase {
  Purchase.customer: Customer
  Purchase.purchasable: Package
  Purchase.expires_at: datetime
  Purchase.status: string
  Purchase.created_at: datetime
}

# -------------------
# Device Token
# -------------------
type DeviceToken {
  DeviceToken.customer: Customer
  DeviceToken.app_version: string
  DeviceToken.platform: string
  DeviceToken.created_at: datetime
}

# -------------------
# Device Authorization
# -------------------
type DeviceAuthorization {
  DeviceAuthorization.customer: Customer
  DeviceAuthorization.type: string
  DeviceAuthorization.last_used_at: datetime
  DeviceAuthorization.created_at: datetime
}

# -------------------
# Payment
# -------------------
type Payment {
  Payment.customer: Customer
  Payment.method: string
  Payment.created_at: datetime
}

# -------------------
# Package
# -------------------
type Package {
  Package.name: string
  Package.purchases: [Purchase]
}

# ===================
# Predicates
# ===================
Customer.date_of_birth: date @index(day) .
Customer.email: string @index(exact) .
Customer.gender: string @index(exact) .
Customer.last_login_at: datetime @index(hour) .
Customer.phone: string @index(exact) .
Customer.created_at: datetime @index(hour) .
Customer.deleted_at: datetime @index(hour) .
Customer.watch_histories: [uid] @reverse .
Customer.purchases: [uid] @reverse .
Customer.device_tokens: [uid] @reverse .
Customer.device_authorizations: [uid] @reverse .
Customer.payments: [uid] @reverse .

WatchHistory.customer: uid @reverse .
WatchHistory.content_id: int @index(int) .
WatchHistory.content_type: string @index(exact) .
WatchHistory.deleted_at: datetime @index(hour) .

Purchase.customer: uid @reverse .
Purchase.purchasable: uid @reverse .
Purchase.expires_at: datetime @index(hour) .
Purchase.status: string @index(exact) .
Purchase.created_at: datetime @index(hour) .

DeviceToken.customer: uid @reverse .
DeviceToken.app_version: string @index(exact) .
DeviceToken.platform: string @index(exact) .
DeviceToken.created_at: datetime @index(hour) .

DeviceAuthorization.customer: uid @reverse .
DeviceAuthorization.type: string @index(exact) .
DeviceAuthorization.last_used_at: datetime @index(hour) .
DeviceAuthorization.created_at: datetime @index(hour) .

Payment.customer: uid @reverse .
Payment.method: string @index(exact) .
Payment.created_at: datetime @index(hour) .

Package.name: string @index(exact) .
Package.purchases: [uid] @reverse .`
}
