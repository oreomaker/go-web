namespace go api

struct BaseResp {
    1: i64 code
    2: string message
}

struct User {
    1: i64 user_id
    2: string username
    3: string avatar
}

struct CreateUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: BaseResp base_resp
}

struct CreateNoteRequest {
    1: string title (api.form="title", api.vd="len($) > 0")
    2: string content (api.form="content", api.vd="len($) > 0")
    3: i64 user_id
}

service ApiService {
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/v2/user/register")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/v2/user/login")
}