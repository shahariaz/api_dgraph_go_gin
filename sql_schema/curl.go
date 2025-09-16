curl -X POST -H "Content-Type: application/json" \
  -d '{
    "query": "query Customers($cutoffDate: datetime, $watchContentType: string, $watchContentIDs: [int], $purchasableID: string, $purchaseStatuses: [string]) { customers(func: type(Customer)) @filter( eq(Customer.deleted_at, null) AND ( ge(Customer.date_of_birth, $cutoffDate) OR anyofterms(Customer.watch_histories.content_type, $watchContentType) OR has(Customer.watch_histories) @filter(eq(WatchHistory.content_id, $watchContentIDs)) ) ) { uid Customer.email Customer.date_of_birth Customer.gender Customer.last_login_at Customer.phone Customer.created_at Customer.watch_histories @filter( eq(WatchHistory.deleted_at, null) AND eq(WatchHistory.content_type, $watchContentType) AND eq(WatchHistory.content_id, $watchContentIDs) ) { WatchHistory.content_id WatchHistory.content_type WatchHistory.deleted_at } Customer.purchases @filter( eq(Purchase.purchasable.Package.name, $purchasableID) OR anyofterms(Purchase.status, $purchaseStatuses) ) { Purchase.status Purchase.expires_at Purchase.created_at Purchase.purchasable { Package.name } } } }",
    "variables": {
      "cutoffDate": "1990-01-01",
      "watchContentType": "Series",
      "watchContentIDs": [219, 225, 232],
      "purchasableID": "Basic Plan",
      "purchaseStatuses": ["active"]
    }
  }' \
  http://localhost:8080/query?timeout=20s