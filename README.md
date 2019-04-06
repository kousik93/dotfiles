# dotfiles

Steps for Mac:

- Install brew
  ```
  $ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
  ```

- Install Iterm2
  ```
  $ brew tap caskroom/cask
  $ brew cask install iterm2
  ```

- Install Misc
  ```
  $ brew install watch && brew install wget
  ```

- Install oh-my-zsh. Copy .zshrc, .oh-my-zsh, .vimrc to the correct place
  ```
  $ cp -r ./.oh-my-zsh ~/
  $ cp ./.vimrc ~/
  $ cp ./.zshrc ~/
  ```

- Install oh-my-zsh autojump ($ brew install autojump)
  ```
  $ brew install autojump
  ```
  - Source zshrc
    ```
    $ source ~/.zshrc
    ```

- Install hack-regular-nerdfont-complete
  ```
  $ brew tap caskroom/fonts
  $ brew cask install font-hack-nerd-font
  ```

- Make zsh default shell
  ```
  chsh -s $(which zsh)
  ```

- Choose to load the Iterm profile from iterm-profile folder in preferences -> general

- Install the Monaco complete font if you want
  - Manually import it into iterm if not already done

- Import the gruvbox-dark.itermcolors to iterm2

- Wget Atom
  ```
  wget -O atom-mac.zip https://atom.io/download/mac
  ```

- Extract and install it

- Move the .atom folder to correct place
  ```
  cr -p ./.atom ~/
  ```


References:
https://luispuerto.net/blog/2018/01/09/iterm2-oh-my-zsh-powerlevel9k-monaco-nerd-complete-font/

https://github.com/ryanoasis/nerd-fonts
