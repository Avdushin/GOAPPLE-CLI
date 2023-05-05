#!/bin/bash

app = "GOAPPLE"

echo -e "\e[0;92$app setup script starting...\e[0m"
# Updating
sudo pacman -Syyyuuu # manjaro
sudo eopkg up -y # solus
sudo dnf update -y # fedora
clear
echo -e "\e[0;92mInstalling programs...\e[0m"
# == installing packages == #

# Manjaro
sudo pacman -Syu i3 i3blocks i3status i3lock polybar kitty krita fish ack vim bottom neofetch flameshot variety feh rofi discord python-pip steam telegram-desktop gcolor3 lxappearance picom flatpak xdg-desktop-portal-gtk awesome-terminal-fonts noto-fonts-emoji noto-fonts bspwm sxhkd --noconfirm
# Solus
sudo eopkg it i3 i3blocks i3status i3-devel i3lock polybar kitty krita xkill fish ack vim bottom neofetch variety feh flameshot rofi discord steam telegram gcolor3 lxappearance picom flatpak xdg-desktop-portal-gtk font-awesome-4 font-awesome-ttf bspwm sxhkd -y 
#Fedora
sudo dnf install \
  https://download1.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm -y && sudo dnf install \
  https://download1.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
 i3 i3-gaps i3blocks i3status i3lock polybar kitty krita fish git ack vim bottom neofetch flameshot variety feh rofi discord steam telegram-desktop picom gcolor3 lxappearance flatpak xdg-desktop-portal-gtk fontawesome-fonts bspwm sxhkd -y --allowerasing --skip-broken
# == Creating directories == #
echo -e "\e[0;92mCreating directories...\e[0m"
sudo mkdir -p ~/.config/i3/
sudo mkdir -p ~/.local/share/Trash/files
# == Unpacking apps (NoiseTorch & Sublime Text) == #
echo -e "\e[0;92mInstalling programs...\e[0m"
sudo tar -C $HOME -xzf src/packages/NoiseTorch_x64.tgz
sudo tar -xf src/packages/sublime_text.tar.xz -C /opt/ && gtk-update-icon-cache && flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
# == Set configs == #
# I3 config
sudo cp -rf src/wms/i3/ ~/.config/
# BSPWM config (bspwmrc, sxhkd, bspwm-polybar-config)
sudo cp -rf src/wms/bspwm/* ~/.config/
# Polybar config
sudo cp -r src/polybar/ ~/.config/
# Kitty terminal config
sudo cp -r src/dots/kitty/ ~/.config/
# neofetch config
sudo cp -rf src/dots/neofetch/ ~/.config/
#sudo flatpak install flathub com.rafaelmardojai.Blanket -y 
sudo cp -r src/dots/.bashrc ~/ 
# Rofi config
sudo cp -rf src/dots/rofi ~/.config/
# == Setting fish-shell by default ==#
# For root and simple $user
sudo chsh -s /usr/bin/fish
chsh -s /usr/bin/fish
# Themes and styles
echo -e "\e[0;92mInstalling themes and fonts...\e[0m"
# set wallpaper
sudo cp -r src/assets/walls/ ~/
# Set icons
git clone https://github.com/PIN3APPLEZZ/tella-icons
sudo cp -r tella-icons/Tela-blue tella-icons/Tela-blue-dark /usr/share/icons/
# Set Themes
git clone https://github.com/PIN3APPLEZZ/themes-for-pin3apple
sudo cp -r themes-for-pin3apple/Solarized-Dark-Blue themes-for-pin3apple/Solarized-Dark-Cyan themes-for-pin3apple/Nordic-darker-standard-buttons /usr/share/themes/
# Set fonts
sudo cp -r src/dots/.fonts ~/
# Get git_token file example
sudo cp -r src/dots/git_token ~/

# BetterLockScreen
wget https://github.com/betterlockscreen/betterlockscreen/archive/refs/heads/main.zip
unzip main.zip
mkdir -p ~/.config/betterlockscreen/
sudo cp /usr/share/doc/betterlockscreen/examples/betterlockscreenrc ~/.config/betterl
ockscreen/
sudo cp betterlockscreen@.service /usr/lib/systemd/system/
systemctl enable betterlockscreen@$USER
betterlockscreen dim -u src/assets/walls/
echo -e "\e[0;92mBetterlockscreen is redy!\e[0m"
echo -e "\e[0;92mDONE!\e[0m"
echo -e "\e[0;91m!\e[0;1;33mYour system will be \e[0;91mREBOOT!\e[0m"
sleep 5
reboot