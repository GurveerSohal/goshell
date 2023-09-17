export const commands = {
    ls : `For each operand that names a file of a type other than directory, ls displays its
    name as well as any requested, associated information.  For each operand that
    names a file of type directory, ls displays the names of files contained within
    that directory, as well as any requested, associated information. <br/> <br/>

    If no operands are given, the contents of the current directory are displayed.  If
    more than one operand is given, non-directory operands are displayed first;
    directory and non-directory operands are sorted separately and in lexicographical
    order.`,

    curl: `curl is a tool for transfering data from or to a server. It supports these
    protocols: DICT, FILE, FTP, FTPS, GOPHER, GOPHERS, HTTP, HTTPS, IMAP, IMAPS,
    LDAP, LDAPS, MQTT, POP3, POP3S, RTMP, RTMPS, RTSP, SCP, SFTP, SMB, SMBS, SMTP,
    SMTPS, TELNET or TFTP. The command is designed to work without user interaction.  <br/> <br/>

    curl offers a busload of useful tricks like proxy support, user authentication,
    FTP upload, HTTP post, SSL connections, cookies, file transfer resume and more.
    As you will see below, the number of features will make your head spin!`,

    cat: `The cat utility reads files sequentially, writing them to the standard output.
    The file operands are processed in command-line order.  If file is a single dash
    (‘-’) or absent, cat reads from the standard input.  If file is a UNIX domain
    socket, cat connects to it and then reads it until EOF.  This complements the UNIX
    domain binding capability available in inetd(8).
    `,

    echo: `The echo utility writes any specified operands, separated by single blank (‘ ’)
    characters and followed by a newline (‘\n’) character, to the standard output.`
  }