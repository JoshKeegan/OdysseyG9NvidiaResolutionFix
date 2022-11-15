# Odyssey G9 Nvidia Resolution Fix
Adds additional resolutions at 240Hz, with adaptive sync, for the Samsung Odyssey G9 when using an Nvidia GPU. 
This allows for games to be played at 16:9 and 21:9 aspect ratios.

## Usage
Run with administrative privileges. You will be required to restart your computer before the new resolutions are available.  
It must be re-run after each GPU driver update (why I wrote this program - doing this manually gets old...).

## Nvidia Control Panel Settings
In order to get these additional resolutions to work with the monitor at 240Hz, the scaling must be handled by the GPU 
so that the monitor still receives a 5120x1440@240Hz signal. This is easily configured within the Nvidia Control Panel.

Under "Display" > "Adjust desktop size and position", select the Odyssey G9.
Then under "Select a scaling mode" select "No Scaling", and under "Perform scaling on" select "GPU".
Apply the settings & select to keep them when prompted.

NVCP settings are retained between driver updates, so this shouldn't need to be repeated.

## Thanks
Thanks to [/u/SpaghettiDuders](https://www.reddit.com/user/spaghettiduders/) for posting 
[his workaround](https://www.reddit.com/r/ultrawidemasterrace/comments/r6r4na/g9_3440x14402560x1439_240hz_gsync_workaround_part/) 
on Reddit. This is just automating his work.

## Licence
[MIT](LICENSE)