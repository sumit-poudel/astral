package main

type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")
const userIdConkextKey = contextKey("userId")
const userNameConkextKey = contextKey("userName")
