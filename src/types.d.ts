declare module "*.svg" {
  const content: string;
  export default content;
}

declare module "*.png" {
  const content: string;
  export default content;
}

declare module "*.jpg" {
  const content: string;
  export default content;
}


export interface Device {
  name: string;
  vendor: string;
  tdp: boolean;
  max_tdp: number;
  min_tdp: number;
  rgb: boolean;
  boost: boolean;
}

export interface Settings {
  app_id: number;
  boost: boolean;
  tdp: boolean;
  tdp_limit: number;
  egpu: boolean;
  debug: boolean;
}

export interface Error {
  error: string;
}

export interface TDPInformation{
  tdp_limit: number;
}