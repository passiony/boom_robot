@echo off
set protoPath=../proto
set outPath=../protogo

rem %cd%/protoc  -I=.  --go_out=.  *.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/login.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/common.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/match.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/home.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/league.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/game.proto
%cd%/protoc  -I=%protoPath%  --go_out=%outPath%  %protoPath%/server.proto

%cd%/gen_cmd.exe -from %protoPath%/msgid.conf -to %outPath%/cmdID.go -name CMDID_2_CMDNAME
