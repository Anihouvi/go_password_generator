
# Go Secure Password generator 
A password generator tools to create or generate a secure password using SHA-256 algorithm by default. 

# Set up your environment

Make sure that you have your Go environment ready and  use the Go command to download all the packages in the .mod file 

# Run the script

The script has a generated executable, you can just run the following:

`./password_generator`

You can also create an alias on macOs to run the script more easily. 
First, you need to make sure the compiled binary is in a directory that is in your `PATH`. If you want to use `/usr/local/bin`, you can move the binary there:

`sudo mv password_generator /usr/local/bin`

Now, open your shell configuration file. If you're using the default Bash shell, the file is` ~/.bashrc or ~/.bash_profile`. If you're using the Zsh shell, the file is `~/.zshrc`. Open the appropriate file in your favorite text editor, or create it if it doesn't exist.

`nano ~/.bashrc`        # for Bash
or
`nano ~/.zshrc `        # for Zsh

Add the following line at the end of the file to create an alias:

`alias passgen='password_generator'`

`source ~/.bashrc`      # for Bash
 or
`source ~/.zshrc`       # for Zsh

Now you should be able to use the passgen alias to run the password_generator script:

`passgen`


#Build up the project from scratch

## Use go get command  to install  packages

`go get <package_name>`

`go run password_generator.go`




## Acknowledgements

 - [All go packages](https://pkg.go.dev/awesome-README-templates)


## 🚀 About Me
I'm a passionate about solving problems and writing interesting codes.Introvert and a Jesus believer. Follow me for more interesting projects. 

➡️ Enter the desired password length: 15
🚏 Using default algorithm: SHA-256
-------------------------------------------------------------------
Length              |Algorithm                     |Hash
-------------------------------------------------------------------
15                  |SHA-256                       |baf80511327c50d
                    |                              |b54c617c5860c6d
                    |                              |4512b7102b139de
                    |                              |7ab654605b441a5
-------------------------------------------------------------------
✅🔐Password Generated
-------------------------------------------------------------------
]z_H&&8"XSaNjM]
-------------------------------------------------------------------
