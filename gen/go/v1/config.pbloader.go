// Code generated by github.com/hypertrace/agent-config/tools/go-generator. DO NOT EDIT.

package v1

import wrappers "google.golang.org/protobuf/types/known/wrapperspb"

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *AgentConfig) loadFromEnv(prefix string, defaultValues *AgentConfig) {
	if x.Opa == nil {
		x.Opa = new(Opa)
	}
	if defaultValues == nil {
		x.Opa.loadFromEnv(prefix+"OPA_", nil)
	} else {
		x.Opa.loadFromEnv(prefix+"OPA_", defaultValues.Opa)
	}

	if x.BlockingConfig == nil {
		x.BlockingConfig = new(BlockingConfig)
	}
	if defaultValues == nil {
		x.BlockingConfig.loadFromEnv(prefix+"BLOCKING_CONFIG_", nil)
	} else {
		x.BlockingConfig.loadFromEnv(prefix+"BLOCKING_CONFIG_", defaultValues.BlockingConfig)
	}

	if val, ok := getBoolEnv(prefix + "DEBUG_LOG"); ok {
		x.DebugLog = &wrappers.BoolValue{Value: val}
	} else if x.DebugLog == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.DebugLog = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.DebugLog != nil {
			x.DebugLog = &wrappers.BoolValue{Value: defaultValues.DebugLog.Value}
		}
	}
	if x.RemoteConfig == nil {
		x.RemoteConfig = new(RemoteConfig)
	}
	if defaultValues == nil {
		x.RemoteConfig.loadFromEnv(prefix+"REMOTE_CONFIG_", nil)
	} else {
		x.RemoteConfig.loadFromEnv(prefix+"REMOTE_CONFIG_", defaultValues.RemoteConfig)
	}

	if x.ApiDiscovery == nil {
		x.ApiDiscovery = new(ApiDiscoveryConfig)
	}
	if defaultValues == nil {
		x.ApiDiscovery.loadFromEnv(prefix+"API_DISCOVERY_", nil)
	} else {
		x.ApiDiscovery.loadFromEnv(prefix+"API_DISCOVERY_", defaultValues.ApiDiscovery)
	}

	if x.Sampling == nil {
		x.Sampling = new(SamplingConfig)
	}
	if defaultValues == nil {
		x.Sampling.loadFromEnv(prefix+"SAMPLING_", nil)
	} else {
		x.Sampling.loadFromEnv(prefix+"SAMPLING_", defaultValues.Sampling)
	}

	if x.Javaagent == nil {
		x.Javaagent = new(Javaagent)
	}
	if defaultValues == nil {
		x.Javaagent.loadFromEnv(prefix+"JAVAAGENT_", nil)
	} else {
		x.Javaagent.loadFromEnv(prefix+"JAVAAGENT_", defaultValues.Javaagent)
	}

	if x.Logging == nil {
		x.Logging = new(LogConfig)
	}
	if defaultValues == nil {
		x.Logging.loadFromEnv(prefix+"LOGGING_", nil)
	} else {
		x.Logging.loadFromEnv(prefix+"LOGGING_", defaultValues.Logging)
	}

	if x.MetricsConfig == nil {
		x.MetricsConfig = new(MetricsConfig)
	}
	if defaultValues == nil {
		x.MetricsConfig.loadFromEnv(prefix+"METRICS_CONFIG_", nil)
	} else {
		x.MetricsConfig.loadFromEnv(prefix+"METRICS_CONFIG_", defaultValues.MetricsConfig)
	}

	if val, ok := getStringEnv(prefix + "ENVIRONMENT"); ok {
		x.Environment = &wrappers.StringValue{Value: val}
	} else if x.Environment == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Environment = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Environment != nil {
			x.Environment = &wrappers.StringValue{Value: defaultValues.Environment.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *Opa) loadFromEnv(prefix string, defaultValues *Opa) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "ENDPOINT"); ok {
		x.Endpoint = &wrappers.StringValue{Value: val}
	} else if x.Endpoint == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Endpoint = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Endpoint != nil {
			x.Endpoint = &wrappers.StringValue{Value: defaultValues.Endpoint.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "POLL_PERIOD_SECONDS"); ok {
		x.PollPeriodSeconds = &wrappers.Int32Value{Value: val}
	} else if x.PollPeriodSeconds == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.PollPeriodSeconds = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.PollPeriodSeconds != nil {
			x.PollPeriodSeconds = &wrappers.Int32Value{Value: defaultValues.PollPeriodSeconds.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "CERT_FILE"); ok {
		x.CertFile = &wrappers.StringValue{Value: val}
	} else if x.CertFile == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.CertFile = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.CertFile != nil {
			x.CertFile = &wrappers.StringValue{Value: defaultValues.CertFile.Value}
		}
	}
	if val, ok := getBoolEnv(prefix + "USE_SECURE_CONNECTION"); ok {
		x.UseSecureConnection = &wrappers.BoolValue{Value: val}
	} else if x.UseSecureConnection == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.UseSecureConnection = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.UseSecureConnection != nil {
			x.UseSecureConnection = &wrappers.BoolValue{Value: defaultValues.UseSecureConnection.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *BlockingConfig) loadFromEnv(prefix string, defaultValues *BlockingConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if val, ok := getBoolEnv(prefix + "DEBUG_LOG"); ok {
		x.DebugLog = &wrappers.BoolValue{Value: val}
	} else if x.DebugLog == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.DebugLog = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.DebugLog != nil {
			x.DebugLog = &wrappers.BoolValue{Value: defaultValues.DebugLog.Value}
		}
	}
	if x.Modsecurity == nil {
		x.Modsecurity = new(ModsecurityConfig)
	}
	if defaultValues == nil {
		x.Modsecurity.loadFromEnv(prefix+"MODSECURITY_", nil)
	} else {
		x.Modsecurity.loadFromEnv(prefix+"MODSECURITY_", defaultValues.Modsecurity)
	}

	if val, ok := getBoolEnv(prefix + "EVALUATE_BODY"); ok {
		x.EvaluateBody = &wrappers.BoolValue{Value: val}
	} else if x.EvaluateBody == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.EvaluateBody = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.EvaluateBody != nil {
			x.EvaluateBody = &wrappers.BoolValue{Value: defaultValues.EvaluateBody.Value}
		}
	}
	if x.RegionBlocking == nil {
		x.RegionBlocking = new(RegionBlockingConfig)
	}
	if defaultValues == nil {
		x.RegionBlocking.loadFromEnv(prefix+"REGION_BLOCKING_", nil)
	} else {
		x.RegionBlocking.loadFromEnv(prefix+"REGION_BLOCKING_", defaultValues.RegionBlocking)
	}

	if x.RemoteConfig == nil {
		x.RemoteConfig = new(RemoteConfig)
	}
	if defaultValues == nil {
		x.RemoteConfig.loadFromEnv(prefix+"REMOTE_CONFIG_", nil)
	} else {
		x.RemoteConfig.loadFromEnv(prefix+"REMOTE_CONFIG_", defaultValues.RemoteConfig)
	}

	if val, ok := getBoolEnv(prefix + "SKIP_INTERNAL_REQUEST"); ok {
		x.SkipInternalRequest = &wrappers.BoolValue{Value: val}
	} else if x.SkipInternalRequest == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.SkipInternalRequest = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.SkipInternalRequest != nil {
			x.SkipInternalRequest = &wrappers.BoolValue{Value: defaultValues.SkipInternalRequest.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "RESPONSE_STATUS_CODE"); ok {
		x.ResponseStatusCode = &wrappers.Int32Value{Value: val}
	} else if x.ResponseStatusCode == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.ResponseStatusCode = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.ResponseStatusCode != nil {
			x.ResponseStatusCode = &wrappers.Int32Value{Value: defaultValues.ResponseStatusCode.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "MAX_RECURSION_DEPTH"); ok {
		x.MaxRecursionDepth = &wrappers.Int32Value{Value: val}
	} else if x.MaxRecursionDepth == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.MaxRecursionDepth = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.MaxRecursionDepth != nil {
			x.MaxRecursionDepth = &wrappers.Int32Value{Value: defaultValues.MaxRecursionDepth.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "RESPONSE_MESSAGE"); ok {
		x.ResponseMessage = &wrappers.StringValue{Value: val}
	} else if x.ResponseMessage == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.ResponseMessage = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.ResponseMessage != nil {
			x.ResponseMessage = &wrappers.StringValue{Value: defaultValues.ResponseMessage.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *ModsecurityConfig) loadFromEnv(prefix string, defaultValues *ModsecurityConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *RegionBlockingConfig) loadFromEnv(prefix string, defaultValues *RegionBlockingConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *RemoteConfig) loadFromEnv(prefix string, defaultValues *RemoteConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "ENDPOINT"); ok {
		x.Endpoint = &wrappers.StringValue{Value: val}
	} else if x.Endpoint == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Endpoint = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Endpoint != nil {
			x.Endpoint = &wrappers.StringValue{Value: defaultValues.Endpoint.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "POLL_PERIOD_SECONDS"); ok {
		x.PollPeriodSeconds = &wrappers.Int32Value{Value: val}
	} else if x.PollPeriodSeconds == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.PollPeriodSeconds = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.PollPeriodSeconds != nil {
			x.PollPeriodSeconds = &wrappers.Int32Value{Value: defaultValues.PollPeriodSeconds.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "CERT_FILE"); ok {
		x.CertFile = &wrappers.StringValue{Value: val}
	} else if x.CertFile == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.CertFile = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.CertFile != nil {
			x.CertFile = &wrappers.StringValue{Value: defaultValues.CertFile.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "GRPC_MAX_CALL_RECV_MSG_SIZE"); ok {
		x.GrpcMaxCallRecvMsgSize = &wrappers.Int32Value{Value: val}
	} else if x.GrpcMaxCallRecvMsgSize == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.GrpcMaxCallRecvMsgSize = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.GrpcMaxCallRecvMsgSize != nil {
			x.GrpcMaxCallRecvMsgSize = &wrappers.Int32Value{Value: defaultValues.GrpcMaxCallRecvMsgSize.Value}
		}
	}
	if val, ok := getBoolEnv(prefix + "USE_SECURE_CONNECTION"); ok {
		x.UseSecureConnection = &wrappers.BoolValue{Value: val}
	} else if x.UseSecureConnection == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.UseSecureConnection = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.UseSecureConnection != nil {
			x.UseSecureConnection = &wrappers.BoolValue{Value: defaultValues.UseSecureConnection.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *ApiDiscoveryConfig) loadFromEnv(prefix string, defaultValues *ApiDiscoveryConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *SamplingConfig) loadFromEnv(prefix string, defaultValues *SamplingConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if x.DefaultRateLimitConfig == nil {
		x.DefaultRateLimitConfig = new(RateLimitConfig)
	}
	if defaultValues == nil {
		x.DefaultRateLimitConfig.loadFromEnv(prefix+"DEFAULT_RATE_LIMIT_CONFIG_", nil)
	} else {
		x.DefaultRateLimitConfig.loadFromEnv(prefix+"DEFAULT_RATE_LIMIT_CONFIG_", defaultValues.DefaultRateLimitConfig)
	}

}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *Javaagent) loadFromEnv(prefix string, defaultValues *Javaagent) {
	if val, ok := getBoolEnv(prefix + "IMPORT_JKS_CERTS"); ok {
		x.ImportJksCerts = &wrappers.BoolValue{Value: val}
	} else if x.ImportJksCerts == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.ImportJksCerts = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.ImportJksCerts != nil {
			x.ImportJksCerts = &wrappers.BoolValue{Value: defaultValues.ImportJksCerts.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *LogConfig) loadFromEnv(prefix string, defaultValues *LogConfig) {
	if rawVal, ok := getStringEnv(prefix + "LOG_MODE"); ok {
		x.LogMode = LogMode(LogMode_value[rawVal])
	} else if x.LogMode == LogMode(0) && defaultValues != nil && defaultValues.LogMode != LogMode(0) {
		x.LogMode = defaultValues.LogMode
	}

	if rawVal, ok := getStringEnv(prefix + "LOG_LEVEL"); ok {
		x.LogLevel = LogLevel(LogLevel_value[rawVal])
	} else if x.LogLevel == LogLevel(0) && defaultValues != nil && defaultValues.LogLevel != LogLevel(0) {
		x.LogLevel = defaultValues.LogLevel
	}

	if x.LogFile == nil {
		x.LogFile = new(LogFileConfig)
	}
	if defaultValues == nil {
		x.LogFile.loadFromEnv(prefix+"LOG_FILE_", nil)
	} else {
		x.LogFile.loadFromEnv(prefix+"LOG_FILE_", defaultValues.LogFile)
	}

}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *LogFileConfig) loadFromEnv(prefix string, defaultValues *LogFileConfig) {
	if val, ok := getInt32Env(prefix + "MAX_FILES"); ok {
		x.MaxFiles = &wrappers.Int32Value{Value: val}
	} else if x.MaxFiles == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.MaxFiles = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.MaxFiles != nil {
			x.MaxFiles = &wrappers.Int32Value{Value: defaultValues.MaxFiles.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "MAX_FILE_SIZE"); ok {
		x.MaxFileSize = &wrappers.Int32Value{Value: val}
	} else if x.MaxFileSize == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.MaxFileSize = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.MaxFileSize != nil {
			x.MaxFileSize = &wrappers.Int32Value{Value: defaultValues.MaxFileSize.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "FILE_PATH"); ok {
		x.FilePath = &wrappers.StringValue{Value: val}
	} else if x.FilePath == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.FilePath = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.FilePath != nil {
			x.FilePath = &wrappers.StringValue{Value: defaultValues.FilePath.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *MetricsLogConfig) loadFromEnv(prefix string, defaultValues *MetricsLogConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "FREQUENCY"); ok {
		x.Frequency = &wrappers.StringValue{Value: val}
	} else if x.Frequency == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Frequency = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Frequency != nil {
			x.Frequency = &wrappers.StringValue{Value: defaultValues.Frequency.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *EndpointMetricsConfig) loadFromEnv(prefix string, defaultValues *EndpointMetricsConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "MAX_ENDPOINTS"); ok {
		x.MaxEndpoints = &wrappers.Int32Value{Value: val}
	} else if x.MaxEndpoints == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.MaxEndpoints = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.MaxEndpoints != nil {
			x.MaxEndpoints = &wrappers.Int32Value{Value: defaultValues.MaxEndpoints.Value}
		}
	}
	if x.Logging == nil {
		x.Logging = new(MetricsLogConfig)
	}
	if defaultValues == nil {
		x.Logging.loadFromEnv(prefix+"LOGGING_", nil)
	} else {
		x.Logging.loadFromEnv(prefix+"LOGGING_", defaultValues.Logging)
	}

}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *MetricsConfig) loadFromEnv(prefix string, defaultValues *MetricsConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if x.EndpointConfig == nil {
		x.EndpointConfig = new(EndpointMetricsConfig)
	}
	if defaultValues == nil {
		x.EndpointConfig.loadFromEnv(prefix+"ENDPOINT_CONFIG_", nil)
	} else {
		x.EndpointConfig.loadFromEnv(prefix+"ENDPOINT_CONFIG_", defaultValues.EndpointConfig)
	}

	if x.Logging == nil {
		x.Logging = new(MetricsLogConfig)
	}
	if defaultValues == nil {
		x.Logging.loadFromEnv(prefix+"LOGGING_", nil)
	} else {
		x.Logging.loadFromEnv(prefix+"LOGGING_", defaultValues.Logging)
	}

}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *RateLimitConfig) loadFromEnv(prefix string, defaultValues *RateLimitConfig) {
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
	if val, ok := getInt64Env(prefix + "MAX_COUNT_GLOBAL"); ok {
		x.MaxCountGlobal = &wrappers.Int64Value{Value: val}
	} else if x.MaxCountGlobal == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.MaxCountGlobal = new(wrappers.Int64Value)
		if defaultValues != nil && defaultValues.MaxCountGlobal != nil {
			x.MaxCountGlobal = &wrappers.Int64Value{Value: defaultValues.MaxCountGlobal.Value}
		}
	}
	if val, ok := getInt64Env(prefix + "MAX_COUNT_PER_ENDPOINT"); ok {
		x.MaxCountPerEndpoint = &wrappers.Int64Value{Value: val}
	} else if x.MaxCountPerEndpoint == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.MaxCountPerEndpoint = new(wrappers.Int64Value)
		if defaultValues != nil && defaultValues.MaxCountPerEndpoint != nil {
			x.MaxCountPerEndpoint = &wrappers.Int64Value{Value: defaultValues.MaxCountPerEndpoint.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "REFRESH_PERIOD"); ok {
		x.RefreshPeriod = &wrappers.StringValue{Value: val}
	} else if x.RefreshPeriod == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.RefreshPeriod = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.RefreshPeriod != nil {
			x.RefreshPeriod = &wrappers.StringValue{Value: defaultValues.RefreshPeriod.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "VALUE_EXPIRATION_PERIOD"); ok {
		x.ValueExpirationPeriod = &wrappers.StringValue{Value: val}
	} else if x.ValueExpirationPeriod == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.ValueExpirationPeriod = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.ValueExpirationPeriod != nil {
			x.ValueExpirationPeriod = &wrappers.StringValue{Value: defaultValues.ValueExpirationPeriod.Value}
		}
	}
	if rawVal, ok := getStringEnv(prefix + "SPAN_TYPE"); ok {
		x.SpanType = SpanType(SpanType_value[rawVal])
	} else if x.SpanType == SpanType(0) && defaultValues != nil && defaultValues.SpanType != SpanType(0) {
		x.SpanType = defaultValues.SpanType
	}

}
