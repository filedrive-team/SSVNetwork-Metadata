create table
    cid_record (
        cid varchar(255) not null,
        data longtext null,
        note varchar(30) null,
        is_sync tinyint(1) null,
        created_at bigint null,
        updated_at bigint null,
        constraint cid_record_cid_uindex unique (cid)
    );

create index cid_record_note_index on cid_record (note);

alter table cid_record add primary key (cid);

create table
    operator_base_info (
        id int unsigned not null comment 'operator id ',
        name varchar(30) null,
        owner_address varchar(42) null,
        public_key longtext null,
        active tinyint(1) null,
        fee longtext null,
        score int unsigned null,
        index_in_owner int unsigned null,
        validator_count int unsigned null,
        timestamp bigint null,
        constraint operator_base_info_id_uindex unique (id)
    );

alter table operator_base_info add primary key (id);

create table
    operator_info (
        id int unsigned not null,
        name varchar(30) null,
        owner_address varchar(42) null,
        location longtext null,
        cloud_provider longtext null,
        eth1_node_client longtext null,
        eth2_node_client longtext null,
        description longtext null,
        website_url longtext null,
        twitter_url longtext null,
        linkedin_url longtext null,
        discord_url longtext null,
        telegram_url longtext null,
        relays_supported longtext NULL,
        mev_bost_enabled int(10) DEFAULT NULL,
        logo longtext null,
        signature longtext null,
        timestamp bigint null,
        cid varchar(255) null,
        constraint operator_info_id_uindex unique (id)
    );

create index operator_info_name_index on operator_info (name);

create index
    operator_info_owner_address_index on operator_info (owner_address);

alter table operator_info add primary key (id);