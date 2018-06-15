# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
export ZSH="/Users/kpadmanabhan/.oh-my-zsh"

# Set name of the theme to load. Optionally, if you set this to "random"
# it'll load a random theme each time that oh-my-zsh is loaded.
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes

#Specifying font before theme
POWERLEVEL9K_MODE='nerdfont-complete'
POWERLEVEL9K_SHOW_CHANGESET=true
POWERLEVEL9K_CHANGESET_HASH_LENGTH=6

ZSH_THEME="powerlevel9k/powerlevel9k"

# Set list of themes to load
# Setting this variable when ZSH_THEME=random
# cause zsh load theme from this variable instead of
# looking in ~/.oh-my-zsh/themes/
# An empty array have no effect
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell" "agnoster" )

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion. Case
# sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# The optional three formats: "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(
  git
  colorize
  dircycle
  lol
  helm
  encode64
  kubectl
  zsh-autosuggestions
  alias-tips 
)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/rsa_id"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"


# Enable autojump
[ -f /usr/local/etc/profile.d/autojump.sh ] && . /usr/local/etc/profile.d/autojump.sh


### Powerlevel9k

#Powerlevel9k base config
#POWERLEVEL9K_PROMPT_ON_NEWLINE=true
POWERLEVEL9K_PROMPT_ADD_NEWLINE=true
#POWERLEVEL9K_RPROMPT_ON_NEWLINE=true
#POWERLEVEL9K_MULTILINE_FIRST_PROMPT_PREFIX="%F{cyan}\u256D\u2500%f"
#POWERLEVEL9K_MULTILINE_LAST_PROMPT_PREFIX="%F{014}\u2570%F{cyan}\uF460%F{073}\uF460%F{109}\uF460%f "
POWERLEVEL9K_SHORTEN_DIR_LENGTH=2



#Powerlevel9k font config

#Powerlevel9k Prompt elements
POWERLEVEL9K_LEFT_PROMPT_ELEMENTS=(os_icon root_indicator user kubecontext vcs dir dir_writable)
POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(go_version background_jobs  battery ram time)
POWERLEVEL9K_BATTERY_VERBOSE=false


#Override gruvbox color overrides
POWERLEVEL9K_KUBECONTEXT_BACKGROUND='124'
POWERLEVEL9K_RAM_BACKGROUND='66'

#Override aliases
alias ll="ls -alh"
alias gits="git status"
alias tmp="cd ~/Desktop/tmp"
alias bit="cd ~/Desktop/bitbucket"
alias desktop="cd ~/Desktop"
alias down="cd ~/Downloads"
alias cls="clear"
alias del="cd /Users/kpadmanabhan/Desktop/go/src/hq-stash.corp.proofpoint.com/del"


# Helper functions to switch clusters
dev () {  export KUBECONFIG=~/.kube/config.dev; }
prestaging () {  export KUBECONFIG=~/.kube/config.prestaging; }
staging () {  export KUBECONFIG=~/.kube/config.staging; }
prod () {  export KUBECONFIG=~/.kube/config.prod; }
integration() { export KUBECONFIG=~/.kube/config.integration; }
dev
source <(kubectl completion zsh)

# Exports
export GOPATH="/Users/kpadmanabhan/Desktop/go"


# Misc functions
busybox_ssh() { echo "kubectl -n run -i --tty busybox --image=yauritux/busybox-curl --restart=Never -- sh  "; }
