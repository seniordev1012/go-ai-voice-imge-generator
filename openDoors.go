package main

func setUpPlatformEnvVars() {
	SetEnvironmentVariable("DB_PASSWORD", "")
	SetEnvironmentVariable("DB_HOST", "")
	SetEnvironmentVariable("DB_USER", "")
	SetEnvironmentVariable("DB_NAME", "")
	SetEnvironmentVariable("OPENAI", openKeys)
	SetEnvironmentVariable("SPEECH_KEY", speechKeys)
	SetEnvironmentVariable("OPENWEATHER", openWeatherKeys)
	SetEnvironmentVariable("ANTHROPIC", anthropicKeys)
}
