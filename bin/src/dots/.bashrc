fish

#noisetorch

if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi

echo " 
#fish_config
if status is-interactive
    # CoMainMenuands to run in interactive sessions can go here
set -g fish_greeting
end " > ~/.config/fish/config.fish

echo " 
function gt
        coMainMenuand cat ~/git_token
end
    
function subl
    coMainMenuand .appz/sublime_text/sublime_text
end" >> ~/.config/fish/functions/fish_prompt.fish