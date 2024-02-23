create table account_info
(
    id          bigint auto_increment comment '主键id'
        primary key,
    nickname    varchar(50)                       not null comment '昵称',
    mobile      varchar(11)                       not null comment '手机号',
    avatar      varchar(1000)                     null comment '头像',
    bg_image    varchar(1000)                     null comment '背景图',
    introduce   varchar(100)                      null comment '介绍',
    sign_info   varchar(100)                      null comment '个人签名',
    gender      enum ('MALE', 'FEMALE', 'SECRET') null comment '性别',
    create_time datetime                          not null comment '创建时间',
    update_time datetime                          not null comment '更新时间',
    revision    bigint default 0                  not null comment '版本号',
    password    varchar(200)                      null comment '密码',
    constraint account_info_pk_2
        unique (mobile)
)
    comment '用户信息表';

create table chat_record
(
    id           bigint auto_increment comment '主键id'
        primary key,
    chat_room_id bigint           not null comment '聊天室id',
    sender_id    bigint           not null comment '发送人ID',
    message      blob             not null comment '消息',
    message_type int              not null comment '消息类型',
    send_time    datetime         not null comment '发送时间',
    canceled     bit default b'0' not null comment '是否撤回 0 否 1 是',
    quote_id     bigint           null comment '引用消息id'
)
    comment '聊天记录';

create table chat_room
(
    id              bigint auto_increment comment '主键id'
        primary key,
    account_ids     text             not null comment '账户id',
    type            int              not null comment '类型 0 单聊 1 群聊',
    name            varchar(20)      null comment '聊天室名称',
    newest_msg      blob             null comment '最新消息',
    newest_msg_time datetime         null comment '最新消息时间',
    newest_msg_type int              null comment '消息类型 0 文本信息 1 语音信息 2 图片信息',
    top             bit default b'0' not null comment '是否置顶',
    create_time     datetime         not null comment '创建时间',
    update_time     datetime         not null comment '更新时间'
)
    comment '聊天室信息';

create table contact_info
(
    id               bigint auto_increment comment '主键id'
        primary key,
    user_id          bigint           not null comment '用户id',
    contact_user_id  bigint           not null comment '通讯录用户id',
    create_time      datetime         not null comment '添加时间',
    stared           bit default b'0' not null comment '是否星标朋友',
    nick_name_remark varchar(200)     null comment '昵称备注'
)
    comment '通讯录信息';


