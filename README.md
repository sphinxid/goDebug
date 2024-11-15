# goDebug - HTTP Header Inspector

A lightweight Go-based debugging tool for inspecting HTTP headers, client IP addresses, and request metadata. This tool is particularly useful for debugging proxy configurations, analyzing request headers, and understanding how your application behaves behind load balancers.

## Features

- 🔒 Token-based authentication for secure access
- 📋 Displays all HTTP request headers
- 🌐 Smart client IP detection (supports X-Forwarded-For and X-Real-IP)
- ⏰ Request timestamp logging
- 🔄 Real-time header inspection
- 🛡️ Basic security measures

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/sphinxid/goDebug
cd goDebug
```

2. Build the application:
```bash
go build
```

3. Run the server:
```bash
./goDebug
```

The server will start on port 9090 by default.

## Usage

Send requests to the server with your authentication token:

```bash
curl "http://localhost:9090/?token=YOUR_TOKEN"
```

### Configuration

The following constants can be modified in the code:

- `PORT`: Server port (default: ":9090")
- `CUSTOM_TOKEN`: Authentication token (default: "Aklj12k3jkasdjflkasdjflkasjf13j")

## Example Output

```plaintext
Request at 2024-11-15 14:30:45.123456 -0700 MST
Accept: */*
User-Agent: curl/7.68.0
X-Forwarded-For: 192.168.1.100
Connection: keep-alive
Client IP: 192.168.1.100
```

## Security Considerations

1. **Token Protection**: Always use a strong, unique token in production
2. **HTTPS**: Consider adding TLS support for production use
3. **Rate Limiting**: Implement rate limiting for production environments
4. **Token Rotation**: Regularly rotate the authentication token

## Use Cases

- Debugging reverse proxy configurations
- Validating load balancer headers
- Testing CDN behavior
- Inspecting client request metadata
- Troubleshooting HTTP header modifications

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Technical Requirements

- Go 1.11 or higher
- No external dependencies
