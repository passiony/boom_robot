package protodef
//############################
//##消息ID=消息名称
//############################
//##req：客户端请求
//##rsp：服务器返回
//##ntf：服务器广播
//############################
//#服务器内部协议
var (
	INTER_ON_CONNECT                                  int32 = 1
	INTER_ON_DIS_CONNECT                              int32 = 2
	INTER_REQ_CONNECTION_ID                           int32 = 3
	INTER_RSP_CONNECTION_ID                           int32 = 4
	INTER_NTF_CONNECTION_ID                           int32 = 5
	W2R_ENTER_ROOM                                    int32 = 6
	R2W_ENTER_ROOM                                    int32 = 7
	W2G_SET_GAME_ID                                   int32 = 8
	W2G_SET_PLAYER_ID                                 int32 = 9
	R2W_EXIT_ROOM                                     int32 = 10
	W2G_KICK_OUT_PLAYER                               int32 = 11
	G2R_UPDATE_PLAYER_GATE                            int32 = 12
	W2G_PRINT_LOG                                     int32 = 13
	W_DAY_CHANGE                                      int32 = 14
	W_UPDATE_TO_DB                                    int32 = 16
	W_FORCE_LEAVE_TEAM                                int32 = 17
)
//#kcp握手
var (
	KCP_SYN                                           int32 = 101
	KCP_ACK                                           int32 = 102
	KCP_FIN                                           int32 = 103
)
//#心跳
var (
	LOGIN_REQ_HEART                                   int32 = 201
	LOGIN_RSP_HEART                                   int32 = 202
)
//#登录
var (
	LOGIN_REQ_LOGIN                                   int32 = 2001
	LOGIN_RSP_LOGIN                                   int32 = 2002
	LOGIN_NTF_KICK_OUT                                int32 = 2003
	LOGIN_NTF_STATE_CHANGE                            int32 = 2004
	LOGIN_NTF_MSG_TIP                                 int32 = 2005
)
//#世界服
//#匹配
var (
	MATCH_REQ_MATCH_GAME                              int32 = 3001
	MATCH_RSP_MATCH_GAME                              int32 = 3002
	MATCH_REQ_CREATE_ROOM                             int32 = 3003
	MATCH_RSP_CREATE_ROOM                             int32 = 3004
	MATCH_REQ_JOIN_ROOM                               int32 = 3005
	MATCH_RSP_JOIN_ROOM                               int32 = 3006
	MATCH_REQ_CANCEL_MATCH                            int32 = 3007
	MATCH_NTF_CANCEL_MATCH                            int32 = 3008
	MATCH_REQ_DELETE_ROOM                             int32 = 3009
	MATCH_NTF_DELETE_ROOM                             int32 = 3010
	MATCH_NTF_ROOM_READY                              int32 = 3011
	MATCH_REQ_MATCH_ACTIVITY                          int32 = 3012
	MATCH_RSP_MATCH_ACTIVITY                          int32 = 3013
	MATCH_REQ_CANCEL_MATCH_ACTIVITY                   int32 = 3014
	MATCH_NTF_CANCEL_MATCH_ACTIVITY                   int32 = 3015
)
var (
	MATCH_REQ_CREATE_TEAM                             int32 = 3021
	MATCH_RSP_CREATE_TEAM                             int32 = 3022
	MATCH_REQ_INVITE_JOIN_TEAM                        int32 = 3023
	MATCH_RSP_INVITE_JOIN_TEAM                        int32 = 3024
	MATCH_REQ_LEAVE_TEAM                              int32 = 3025
	MATCH_NTF_LEAVE_TEAM                              int32 = 3026
	MATCH_REQ_SEARCH_PALYERS                          int32 = 3027
	MATCH_NTF_SEARCH_PLAYERS                          int32 = 3028
	MATCH_REQ_MATCH_TEAM_GAME                         int32 = 3029
	MATCH_NTF_MATCH_TEAM_GAME                         int32 = 3030
	MATCH_REQ_BREAK_TEAM                              int32 = 3031
	MATCH_NTF_BREAK_TEAM                              int32 = 3032
	MATCH_REQ_TEAM_MEMBER_READY                       int32 = 3033
	MATCH_NTF_TEAM_MEMBER_READY                       int32 = 3034
	MATCH_NTF_MY_TEAM                                 int32 = 3035
	MATCH_REQ_AGREE_JOIN_TEAM                         int32 = 3036
	MATCH_REQ_SELECT_TEAM_TALENTID                    int32 = 3037
	MATCH_RSP_SELECT_TEAM_TALENTID                    int32 = 3038
	MATCH_RSP_AGREE_JOIN_TEAM                         int32 = 3039
	MATCH_REQ_MY_TEAM                                 int32 = 3040
	MATCH_REQ_UPDATE_TEAM                             int32 = 3041
//# 通知工会内在线玩家 , 有人邀请玩家加入组队
	MATCH_NTF_ANYONE_JOIN_TEAM                        int32 = 3044
)
//#大厅
var (
	HOME_REQ_SAVE_ARCHIVE                             int32 = 3101
	HOME_REQ_PLAYER_INFO                              int32 = 3102
	HOME_NTF_PLAYER_INFO                              int32 = 3103
	HOME_REQ_SET_PLAYER_INFO                          int32 = 3104
	HOME_RSP_SET_PLAYER_INFO                          int32 = 3105
	HOME_REQ_IAP_VERIFY                               int32 = 3109
	HOME_REQ_RANK_DATA                                int32 = 3110
	HOME_REQ_MY_RANK_DATA                             int32 = 3111
	HOME_REQ_SEASON_DATA                              int32 = 3112
	HOME_REQ_RESET_SEASON                             int32 = 3113
	HOME_REQ_ACTIVE_DATA                              int32 = 3114
	HOME_REQ_SAVE_ACTIVE_INFO                         int32 = 3115
	HOME_REQ_LUCKYWHEEL_DATA                          int32 = 3116
	HOME_REQ_RAFFLE_LUCKYWHEEL                        int32 = 3117
	HOME_REQ_DAILYSIGN_DATA                           int32 = 3118
	HOME_REQ_SIGN_IN                                  int32 = 3119
	HOME_REQ_VIDEO_LIST                               int32 = 3120
	HOME_REQ_VIDEO_PLAY                               int32 = 3121
	HOME_REQ_VIDEO_REWARD                             int32 = 3122
	HOME_REQ_PROTECT_WIN_STREAK                       int32 = 3123
	HOME_REQ_FRIEND_REWARD_DATA                       int32 = 3124
	HOME_REQ_SAVE_FRIEND_REWARD                       int32 = 3125
	HOME_REQ_LOTTERY_FRIEND_REWARD                    int32 = 3126
	HOME_REQ_MASTER_RANK_DATA                         int32 = 3127
	HOME_REQ_MASTER_MY_RANK                           int32 = 3128
	HOME_REQ_MASTER_CLEAN_OLD                         int32 = 3129
	HOME_REQ_DIAMOND_CONSUME                          int32 = 3130
	HOME_REQ_UPLOAD_FCM_TOKEN                         int32 = 3131
)
//#小队
var (
	LEAGUE_REQ_CREATE_LEAGUE                          int32 = 3201
	LEAGUE_NTF_CREATE_LEAGUE                          int32 = 3202
	LEAGUE_REQ_MY_LEAGUE                              int32 = 3203
	LEAGUE_NTF_MY_LEAGUE                              int32 = 3204
	LEAGUE_REQ_DISMISS_LEAGUE                         int32 = 3205
	LEAGUE_NTF_DISMISS_LEAGUE                         int32 = 3206
	LEAGUE_REQ_INVITE_OTHER_JOIN_LEAGUE               int32 = 3207
	LEAGUE_NTF_INVITE_OTHER_JOIN_LEAGUE               int32 = 3208
	LEAGUE_NTF_INVITE_ME_JOIN_LEAGUE                  int32 = 3209
	LEAGUE_REQ_JOIN_LEAGUE                            int32 = 3210
	LEAGUE_NTF_JOIN_LEAGUE                            int32 = 3211
	LEAGUE_REQ_LEAVE_LEAGUE                           int32 = 3212
	LEAGUE_NTF_LEAVE_LEAGUE                           int32 = 3213
	LEAGUE_REQ_LEAGUE_LIST                            int32 = 3214
	LEAGUE_NTF_LEAGUE_LIST                            int32 = 3215
	LEAGUE_REQ_INVITE_PLAYER_LIST                     int32 = 3216
	LEAGUE_NTF_INVITE_PLAYER_LIST                     int32 = 3217
	LEAGUE_REQ_MAIL_LIST                              int32 = 3218
	LEAGUE_NTF_MAIL                                   int32 = 3219
	LEAGUE_REQ_DELETE_MAIL                            int32 = 3220
	LEAGUE_RSP_MAIL_LIST                              int32 = 3221
)
//#游戏服
var (
	GAME_REQ_GAME_READY                               int32 = 4001
	GAME_NTF_GAME_START                               int32 = 4002
	GAME_REQ_GAME_OVER                                int32 = 4003
	GAME_NTF_GAME_OVER                                int32 = 4004
	GAME_REQ_SURRENDER                                int32 = 4005
	GAME_REQ_ROUND_CHANGE                             int32 = 4006
	GAME_NTF_ROUND_CHANGE                             int32 = 4007
	GAME_REQ_ROUND_END                                int32 = 4008
	GAME_REQ_ACTION_END                               int32 = 4010
	GAME_NTF_ACTION_END                               int32 = 4011
	GAME_NTF_SYNC_FRAME                               int32 = 4012
	GAME_REQ_GAME_RECONNECT                           int32 = 4013
	GAME_NTF_GAME_RECONNECT                           int32 = 4014
	GAME_REQ_RECONNECT_FINISH                         int32 = 4015
	GAME_NTF_RECONNECT_FINISH                         int32 = 4016
	GAME_NTF_MAP_EVENT                                int32 = 4017
	GAME_NTF_ACTIVITY_GAME_OVER                       int32 = 4018
	GAME_REQ_ROOM_READY                               int32 = 4020
	GAME_REQ_REFRESH_CARDS                            int32 = 4021
	GAME_REQ_REFRESH_HP                               int32 = 4022
	GAME_REQ_REFRESH_MAP_EVENT                        int32 = 4023
	GAME_REQ_ACTION_START                             int32 = 4024
	GAME_NTF_ACTION_START                             int32 = 4025
	GAME_REQ_ADD_SUMMON                               int32 = 4026
	GAME_REQ_BULLET_HIT_ENTITY                        int32 = 4027
	GAME_NTF_COMBO_INSERT                             int32 = 4028
	GAME_NTF_COMBO_JUDGE                              int32 = 4029
	GAME_REQ_DEL_SUMMON                               int32 = 4030
)
var (
	GAME_REQ_SELECT_CARDS                             int32 = 4101
	GAME_NTF_SELECT_CARDS                             int32 = 4102
	GAME_REQ_SELECT_ROLE                              int32 = 4103
	GAME_NTF_SELECT_ROLE                              int32 = 4104
	GAME_REQ_DESELECT_ROLE                            int32 = 4105
	GAME_NTF_DESELECT_ROLE                            int32 = 4106
	GAME_REQ_CHANGE_ROLE                              int32 = 4109
	GAME_NTF_CHANGE_ROLE                              int32 = 4110
	GAME_REQ_ADD_STAY_BULLET                          int32 = 4113
	GAME_REQ_DEL_STAY_BULLET                          int32 = 4114
	GAME_REQ_SYNC_RANDOM                              int32 = 4115
	GAME_NTF_SYNC_RANDOM                              int32 = 4116
	GAME_REQ_ADD_BUFF                                 int32 = 4117
	GAME_REQ_DEL_BUFF                                 int32 = 4118
	GAME_REQ_DAZZLING_ACTION                          int32 = 4119
	GAME_NTF_DAZZLING_ACTION                          int32 = 4120
	GAME_REQ_PLAYER_DETAIL                            int32 = 4121
)
var (
	GAME_REQ_FINGER_MOVE                              int32 = 4201
	GAME_NTF_FINGER_MOVE                              int32 = 4202
	GAME_REQ_SHOOT_BULLET                             int32 = 4203
	GAME_NTF_SHOOT_BULLET                             int32 = 4204
	GAME_REQ_BULLET_TIME                              int32 = 4205
	GAME_NTF_BULLET_TIME                              int32 = 4206
	GAME_REQ_RETURN_BULLET                            int32 = 4207
	GAME_NTF_RETURN_BULLET                            int32 = 4208
	GAME_REQ_REVIVAL_ROLE                             int32 = 4209
	GAME_REQ_RANDOMPOSITION_BULLET                    int32 = 4211
	GAME_NTF_RANDOMPOSITION_BULLET                    int32 = 4212
)
var (
	GAME_REQ_EMOJI                                    int32 = 4301
	GAME_NTF_EMOJI                                    int32 = 4302
	GAME_REQ_ADD_MP                                   int32 = 4303
	GAME_NTF_ADD_MP                                   int32 = 4304
	GAME_REQ_MARK                                     int32 = 4305
	GAME_NTF_MARK                                     int32 = 4306
)
var (
	GAME_REQ_AI_MESSAGE                               int32 = 5000
)
var CMDID_2_CMDNAME map[int32]string
func init() {
	CMDID_2_CMDNAME = make(map[int32]string)
	CMDID_2_CMDNAME[4028] = "game_ntf_combo_insert"
	CMDID_2_CMDNAME[4101] = "game_req_select_cards"
	CMDID_2_CMDNAME[4114] = "game_req_del_stay_bullet"
	CMDID_2_CMDNAME[5000] = "game_req_ai_message"
	CMDID_2_CMDNAME[3040] = "match_req_my_team"
	CMDID_2_CMDNAME[3117] = "home_req_raffle_luckywheel   "
	CMDID_2_CMDNAME[3215] = "league_ntf_league_list"
	CMDID_2_CMDNAME[4004] = "game_ntf_game_over"
	CMDID_2_CMDNAME[3032] = "match_ntf_break_team"
	CMDID_2_CMDNAME[3128] = "home_req_master_my_rank  "
	CMDID_2_CMDNAME[3031] = "match_req_break_team"
	CMDID_2_CMDNAME[3205] = "league_req_dismiss_league"
	CMDID_2_CMDNAME[3209] = "league_ntf_invite_me_join_league"
	CMDID_2_CMDNAME[4016] = "game_ntf_reconnect_finish"
	CMDID_2_CMDNAME[4024] = "game_req_action_start"
	CMDID_2_CMDNAME[10] = "r2w_exit_room"
	CMDID_2_CMDNAME[3029] = "match_req_match_team_game"
	CMDID_2_CMDNAME[4304] = "game_ntf_add_mp"
	CMDID_2_CMDNAME[3021] = "match_req_create_team"
	CMDID_2_CMDNAME[4303] = "game_req_add_mp"
	CMDID_2_CMDNAME[3116] = "home_req_luckywheel_data "
	CMDID_2_CMDNAME[4113] = "game_req_add_stay_bullet"
	CMDID_2_CMDNAME[4302] = "game_ntf_emoji"
	CMDID_2_CMDNAME[14] = "w_day_change"
	CMDID_2_CMDNAME[2002] = "login_rsp_login   "
	CMDID_2_CMDNAME[3207] = "league_req_invite_other_join_league"
	CMDID_2_CMDNAME[4110] = "game_ntf_change_role"
	CMDID_2_CMDNAME[4202] = "game_ntf_finger_move"
	CMDID_2_CMDNAME[4305] = "game_req_mark"
	CMDID_2_CMDNAME[5] = "inter_ntf_connection_id"
	CMDID_2_CMDNAME[13] = "w2g_print_log"
	CMDID_2_CMDNAME[3220] = "league_req_delete_mail"
	CMDID_2_CMDNAME[4109] = "game_req_change_role"
	CMDID_2_CMDNAME[202] = "login_rsp_heart"
	CMDID_2_CMDNAME[3125] = "home_req_save_friend_reward  "
	CMDID_2_CMDNAME[3113] = "home_req_reset_season"
	CMDID_2_CMDNAME[3129] = "home_req_master_clean_old"
	CMDID_2_CMDNAME[4020] = "game_req_room_ready"
	CMDID_2_CMDNAME[4030] = "game_req_del_summon"
	CMDID_2_CMDNAME[4104] = "game_ntf_select_role"
	CMDID_2_CMDNAME[4201] = "game_req_finger_move"
	CMDID_2_CMDNAME[3022] = "match_rsp_create_team"
	CMDID_2_CMDNAME[3023] = "match_req_invite_join_team"
	CMDID_2_CMDNAME[4203] = "game_req_shoot_bullet"
	CMDID_2_CMDNAME[4023] = "game_req_refresh_map_event"
	CMDID_2_CMDNAME[17] = "w_force_leave_team"
	CMDID_2_CMDNAME[3026] = "match_ntf_leave_team"
	CMDID_2_CMDNAME[3201] = "league_req_create_league"
	CMDID_2_CMDNAME[4001] = "game_req_game_ready"
	CMDID_2_CMDNAME[4008] = "game_req_round_end"
	CMDID_2_CMDNAME[4015] = "game_req_reconnect_finish"
	CMDID_2_CMDNAME[4116] = "game_ntf_sync_random"
	CMDID_2_CMDNAME[3003] = "match_req_create_room"
	CMDID_2_CMDNAME[3105] = "home_rsp_set_player_info"
	CMDID_2_CMDNAME[3024] = "match_rsp_invite_join_team"
	CMDID_2_CMDNAME[3038] = "match_rsp_select_team_talentId"
	CMDID_2_CMDNAME[3101] = "home_req_save_archive"
	CMDID_2_CMDNAME[3109] = "home_req_iap_verify  "
	CMDID_2_CMDNAME[3124] = "home_req_friend_reward_data  "
	CMDID_2_CMDNAME[3127] = "home_req_master_rank_data"
	CMDID_2_CMDNAME[4] = "inter_rsp_connection_id"
	CMDID_2_CMDNAME[2001] = "login_req_login   "
	CMDID_2_CMDNAME[4209] = "game_req_revival_role"
	CMDID_2_CMDNAME[4306] = "game_ntf_mark"
	CMDID_2_CMDNAME[3216] = "league_req_invite_player_list"
	CMDID_2_CMDNAME[4013] = "game_req_game_reconnect"
	CMDID_2_CMDNAME[3002] = "match_rsp_match_game"
	CMDID_2_CMDNAME[3035] = "match_ntf_my_team"
	CMDID_2_CMDNAME[3210] = "league_req_join_league"
	CMDID_2_CMDNAME[3221] = "league_rsp_mail_list"
	CMDID_2_CMDNAME[4301] = "game_req_emoji"
	CMDID_2_CMDNAME[3] = "inter_req_connection_id"
	CMDID_2_CMDNAME[2003] = "login_ntf_kick_out"
	CMDID_2_CMDNAME[4005] = "game_req_surrender"
	CMDID_2_CMDNAME[4105] = "game_req_deselect_role"
	CMDID_2_CMDNAME[4204] = "game_ntf_shoot_bullet"
	CMDID_2_CMDNAME[4205] = "game_req_bullet_time"
	CMDID_2_CMDNAME[4208] = "game_ntf_return_bullet"
	CMDID_2_CMDNAME[201] = "login_req_heart"
	CMDID_2_CMDNAME[3218] = "league_req_mail_list"
	CMDID_2_CMDNAME[8] = "w2g_set_game_id"
	CMDID_2_CMDNAME[3009] = "match_req_delete_room"
	CMDID_2_CMDNAME[3041] = "match_req_update_team"
	CMDID_2_CMDNAME[3121] = "home_req_video_play  "
	CMDID_2_CMDNAME[3123] = "home_req_protect_win_streak  "
	CMDID_2_CMDNAME[4026] = "game_req_add_summon"
	CMDID_2_CMDNAME[4117] = "game_req_add_buff"
	CMDID_2_CMDNAME[4211] = "game_req_randomPosition_bullet"
	CMDID_2_CMDNAME[103] = "kcp_fin"
	CMDID_2_CMDNAME[2004] = "login_ntf_state_change"
	CMDID_2_CMDNAME[3027] = "match_req_search_palyers"
	CMDID_2_CMDNAME[3014] = "match_req_cancel_match_activity"
	CMDID_2_CMDNAME[3015] = "match_ntf_cancel_match_activity"
	CMDID_2_CMDNAME[3110] = "home_req_rank_data   "
	CMDID_2_CMDNAME[3112] = "home_req_season_data "
	CMDID_2_CMDNAME[4002] = "game_ntf_game_start"
	CMDID_2_CMDNAME[4007] = "game_ntf_round_change"
	CMDID_2_CMDNAME[2] = "inter_on_dis_connect"
	CMDID_2_CMDNAME[3103] = "home_ntf_player_info "
	CMDID_2_CMDNAME[3008] = "match_ntf_cancel_match"
	CMDID_2_CMDNAME[3013] = "match_rsp_match_activity"
	CMDID_2_CMDNAME[3217] = "league_ntf_invite_player_list"
	CMDID_2_CMDNAME[4003] = "game_req_game_over"
	CMDID_2_CMDNAME[4014] = "game_ntf_game_reconnect"
	CMDID_2_CMDNAME[4102] = "game_ntf_select_cards"
	CMDID_2_CMDNAME[11] = "w2g_kick_out_player"
	CMDID_2_CMDNAME[3007] = "match_req_cancel_match"
	CMDID_2_CMDNAME[4212] = "game_ntf_randomPosition_bullet"
	CMDID_2_CMDNAME[101] = "kcp_syn"
	CMDID_2_CMDNAME[3202] = "league_ntf_create_league"
	CMDID_2_CMDNAME[3036] = "match_req_agree_join_team"
	CMDID_2_CMDNAME[4025] = "game_ntf_action_start"
	CMDID_2_CMDNAME[4103] = "game_req_select_role"
	CMDID_2_CMDNAME[3006] = "match_rsp_join_room"
	CMDID_2_CMDNAME[3131] = "home_req_upload_fcm_token"
	CMDID_2_CMDNAME[3111] = "home_req_my_rank_data"
	CMDID_2_CMDNAME[3122] = "home_req_video_reward"
	CMDID_2_CMDNAME[4011] = "game_ntf_action_end"
	CMDID_2_CMDNAME[4012] = "game_ntf_sync_frame"
	CMDID_2_CMDNAME[4017] = "game_ntf_map_event"
	CMDID_2_CMDNAME[4115] = "game_req_sync_random"
	CMDID_2_CMDNAME[3044] = "match_ntf_anyone_join_team"
	CMDID_2_CMDNAME[3102] = "home_req_player_info "
	CMDID_2_CMDNAME[4119] = "game_req_dazzling_action"
	CMDID_2_CMDNAME[4121] = "game_req_player_detail"
	CMDID_2_CMDNAME[3010] = "match_ntf_delete_room"
	CMDID_2_CMDNAME[3034] = "match_ntf_team_member_ready"
	CMDID_2_CMDNAME[3114] = "home_req_active_data "
	CMDID_2_CMDNAME[3126] = "home_req_lottery_friend_reward   "
	CMDID_2_CMDNAME[3130] = "home_req_diamond_consume "
	CMDID_2_CMDNAME[3203] = "league_req_my_league"
	CMDID_2_CMDNAME[1] = "inter_on_connect"
	CMDID_2_CMDNAME[3005] = "match_req_join_room"
	CMDID_2_CMDNAME[3213] = "league_ntf_leave_league"
	CMDID_2_CMDNAME[4018] = "game_ntf_activity_game_over"
	CMDID_2_CMDNAME[3208] = "league_ntf_invite_other_join_league"
	CMDID_2_CMDNAME[4022] = "game_req_refresh_hp"
	CMDID_2_CMDNAME[4106] = "game_ntf_deselect_role"
	CMDID_2_CMDNAME[4206] = "game_ntf_bullet_time"
	CMDID_2_CMDNAME[4207] = "game_req_return_bullet"
	CMDID_2_CMDNAME[3039] = "match_rsp_agree_join_team"
	CMDID_2_CMDNAME[3206] = "league_ntf_dismiss_league"
	CMDID_2_CMDNAME[3001] = "match_req_match_game"
	CMDID_2_CMDNAME[3033] = "match_req_team_member_ready"
	CMDID_2_CMDNAME[3037] = "match_req_select_team_talentId"
	CMDID_2_CMDNAME[4021] = "game_req_refresh_cards"
	CMDID_2_CMDNAME[7] = "r2w_enter_room"
	CMDID_2_CMDNAME[16] = "w_update_to_db"
	CMDID_2_CMDNAME[3212] = "league_req_leave_league"
	CMDID_2_CMDNAME[2005] = "login_ntf_msg_tip"
	CMDID_2_CMDNAME[3204] = "league_ntf_my_league"
	CMDID_2_CMDNAME[3104] = "home_req_set_player_info"
	CMDID_2_CMDNAME[3219] = "league_ntf_mail"
	CMDID_2_CMDNAME[4029] = "game_ntf_combo_judge"
	CMDID_2_CMDNAME[3011] = "match_ntf_room_ready"
	CMDID_2_CMDNAME[3028] = "match_ntf_search_players"
	CMDID_2_CMDNAME[3115] = "home_req_save_active_info"
	CMDID_2_CMDNAME[4118] = "game_req_del_buff"
	CMDID_2_CMDNAME[4120] = "game_ntf_dazzling_action"
	CMDID_2_CMDNAME[6] = "w2r_enter_room"
	CMDID_2_CMDNAME[3012] = "match_req_match_activity"
	CMDID_2_CMDNAME[9] = "w2g_set_player_id"
	CMDID_2_CMDNAME[4006] = "game_req_round_change"
	CMDID_2_CMDNAME[102] = "kcp_ack"
	CMDID_2_CMDNAME[3211] = "league_ntf_join_league"
	CMDID_2_CMDNAME[3025] = "match_req_leave_team"
	CMDID_2_CMDNAME[3120] = "home_req_video_list  "
	CMDID_2_CMDNAME[3214] = "league_req_league_list"
	CMDID_2_CMDNAME[4010] = "game_req_action_end"
	CMDID_2_CMDNAME[12] = "g2r_update_player_gate"
	CMDID_2_CMDNAME[3004] = "match_rsp_create_room"
	CMDID_2_CMDNAME[3119] = "home_req_sign_in "
	CMDID_2_CMDNAME[4027] = "game_req_bullet_hit_entity"
	CMDID_2_CMDNAME[3030] = "match_ntf_match_team_game"
	CMDID_2_CMDNAME[3118] = "home_req_dailysign_data  "
}
