import { Device, Settings, TDPInformation } from "./types";

const PORT: number = 33238;
const HOST: string = `http://localhost:${PORT}/api`;

export async function getDevice(): Promise<Device> {
    let res = await fetch(`${HOST}/device`);
    return await res.json();
}

export async function getSettings(appId:number): Promise<Settings> {
  let res = await fetch(`${HOST}/settings/${appId}`);
    return await res.json();
}

export async function setSettings(settings:Settings): Promise<Settings> {
  let res = await fetch(`${HOST}/settings/${settings.app_id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(settings),
  });
  return await res.json();
}

export async function getTDPInformation(): Promise<TDPInformation> {
  let res = await fetch(`${HOST}/tdp-information`);
    return await res.json();
}