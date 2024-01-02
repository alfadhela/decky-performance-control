import os
import pathlib
import asyncio
import subprocess

PLUGIN_DIR = str(pathlib.Path(__file__).parent.resolve())

class Plugin:
    backend_proc = None
    # Asyncio-compatible long-running code, executed in a task when the plugin is loaded
    async def _main(self):
        self.backend_proc = subprocess.Popen([PLUGIN_DIR + "/bin/backend"])
        while True:
            await asyncio.sleep(1)
    
    # Make sure to shutdown backend
    async def _unload(self):
        if self.backend_proc is not None:
            self.backend_proc.terminate()
            try:
                self.backend_proc.wait(timeout=2)
            except:
                self.backend_proc.kill()
            self.backend_proc = None
