# yallyscreen [![GitHub release]](https://github.com/Pyxis66/yallyscreen/releases) [![license]]()

_yallyscreen_ is a LCD touch interface for your OctoPrint server.  It is based on GTK+3 and allows you to control your 3D Printer using a [LCD touch screen](https://amzn.to/2L8cRkR), a [Raspberry Pi](https://amzn.to/39LPvvF), and [OctoPrint](https://octoprint.org/).  It's an _X application_ that's executed directly in the X Server without a window manager or browser, and operates as a frontend for OctoPrint.

------------
## Installation

### Dependencies

yallyscreen is based on [GoLang](https://golang.org).  GoLang applications are usually dependency-less, but in this case [GTK+3](https://developer.gnome.org/gtk3/3.0/gtk.html) is used, and the GTK+3 libraries are required in order to run.  Be sure that GTK+3 is installed and is the only graphical environment that's been installed.

If you are using `Raspbian` or any other `Debian` based distribution, required packages can be installed using:

```sh
sudo apt-get install libgtk-3-0 xserver-xorg xinit x11-xserver-utils
```

You will also need to set up the video drivers for the display you are using.  Installation and configuration of the drivers is usually specific to the display you are using, and you will need to contact the manufacturer for instructions.  To help you set up your system and display, a setup guide ([Setting Up yallyscreen and Your Display](https://github.com/Pyxis66/yallyscreen/wiki/Setting-Up-yallyscreen-and-Your-Display)) is available in the wiki.



### Install From Source

The compilation and packaging tasks are managed by the [`Makefile`](Makefile) and backed on [Docker](Dockerfile).  Docker is used to avoid installing any other dependencies since all the operations are done inside of the container.

If you need to install docker inside `Raspbian` or any other linux distribution just run:

```sh
curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh
```

You can read more about this at [`docker-install`](https://github.com/docker/docker-install)

To compile the project (assuming that you already cloned this repository), just execute the `build` target.  This will generate all the binaries and debian packages in the `build` folder:

```sh
make build
```

The default build is for the STRETCH release of debian, but BUSTER and JESSIE are also possible.  To build one of these targets, you just have to specify the package during make.
Example for BUSTER:
```sh
make build DEBIAN_PACKAGES=BUSTER
ls -1 build/
```

If you are using `Raspbian` you can install any of the `.deb` generated packages.  If not, just use the compiled binary.

------------
## Configuration

### Basic Configuration

------------
## Menu Configuration

### Custom Controls and Commands

Custom [controls](http://docs.octoprint.org/en/master/configuration/config_yaml.html#controls) to execute GCODE instructions and [commands](http://docs.octoprint.org/en/master/configuration/config_yaml.html#system) to execute shell commands can be defined in the `config.yaml` file.

The controls are limit to static controls without `inputs`.

------------
## License

GNU Affero General Public License v3.0, see [LICENSE](LICENSE)

This project is a hard fork from [Octoprint-TFT](https://github.com/mcuadros/OctoPrint-TFT) created by [@mcuadros](https://github.com/mcuadros/OctoPrint-TFT)
