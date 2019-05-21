package cache

import (
  "edgex-club/model"
)

//{Token:User}
var TokenCache = make(map[string]model.User, 20)
