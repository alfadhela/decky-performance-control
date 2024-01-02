import {
  Field,
  PanelSection,
  PanelSectionRow,
  ServerAPI,
  SliderField,
  ToggleField,
  definePlugin,
  staticClasses,
} from "decky-frontend-lib";
import { useEffect, useState, VFC } from "react";
import { FaTachometerAlt } from "react-icons/fa";

import { Device, Settings, TDPInformation } from "./types";
import * as backend from "./backend";

const Content: VFC<{ serverAPI: ServerAPI }> = () => {
  const [settings, setSettings] = useState<Settings>({app_id: 0,boost: false, tdp: false, tdp_limit: 10, egpu: false, debug: true})
  const [device, setDevice] = useState<Device>({name: "test", vendor: "", tdp:false, min_tdp: 5, max_tdp: 30, rgb: false, boost: false});
  const [tdpInfo, setTDPInfo] = useState<TDPInformation>({tdp_limit:10});

  useEffect(() => {
    backend.getDevice()
      .then((device) => { setDevice(device) });
  }, []);

  useEffect(() => {
    backend.getSettings(0)
      .then((settings) => { setSettings(settings) });
  }, [])

  useEffect(() => {
    backend.getTDPInformation()
      .then((tdpInfo) => { setTDPInfo(tdpInfo) })
  },[]);

  return (
    <PanelSection title="CPU">
      <Field
        label={device.vendor + " " + device.name +" ("+tdpInfo.tdp_limit+") Watts"}
      />
      <PanelSectionRow>
        <ToggleField 
          label="Boost" 
          description="Increase maximum CPU frequency"
          disabled={!device.boost}
          onChange={(value) => {
            settings.boost = value;
            backend.setSettings(settings);
          }}
          checked={settings.boost}/>
      </PanelSectionRow>
      <PanelSectionRow>
        <ToggleField 
          label="TDP Limit" 
          description="Limit CPU for less total power"
          disabled={!device.tdp}
          onChange={(value) => {
            let newSettings = {app_id: settings.app_id, boost: settings.boost, tdp: value, tdp_limit: settings.tdp_limit, egpu: settings.egpu, debug:settings.debug};
            setSettings(newSettings);
            backend.setSettings(newSettings);
          }}
          checked={settings.tdp}/>
          {settings.tdp && <SliderField
            label="Watts"
            showValue={true}
            max={device.max_tdp}
            min={device.min_tdp}
            step={1}
            value={tdpInfo.tdp_limit}
            onChange={(value) => {
              let newSettings = {app_id: settings.app_id, boost: settings.boost, tdp: settings.tdp, tdp_limit: value, egpu: settings.egpu, debug:settings.debug};
              setSettings(newSettings);
              setTDPInfo({tdp_limit: value});
              backend.setSettings(newSettings);
            }}
            />
          }
      </PanelSectionRow>
    </PanelSection>
  );
};

export default definePlugin((serverApi: ServerAPI) => {
  return {
    title: <div className={staticClasses.Title}>Performance</div>,
    content: <Content serverAPI={serverApi} />,
    icon: <FaTachometerAlt />,
  };
});