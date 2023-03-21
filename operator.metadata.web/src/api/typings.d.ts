export interface OperatorData {
  cid: string;
  infos: OperatorInfo[];
  total: number;
}

export interface OperatorInfo {
  cid: string;
  description: string;
  eth1_node_client: string;
  eth2_node_client: string;
  linkedin_url: string;
  location: string;
  logo: string;
  name: string;
  operator_id: number;
  owner_address: string;
  cloud_provider: string;
  signature: string;
  timestamp: number;
  twitter_url: string;
  website_url: string;
  discord_url: string;
  telegram_url: string;
  mev_bost_enabled: boolean;
  relays_supported: Array;
}
