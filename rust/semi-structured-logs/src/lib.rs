/// various log levels
#[derive(Clone, PartialEq, Eq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
    Debug,
}

/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
    return if level == LogLevel::Info {
        info(message).to_string()
    } else if level == LogLevel::Error {
        error(message).to_string()
    } else if level == LogLevel::Warning {
        warn(message).to_string()
    } else if level == LogLevel::Debug {
        debug(message).to_string()
    } else {
        "".to_string()
    };
}

pub fn info(message: &str) -> String {
    return format!("[INFO]: {}", message);
}

pub fn warn(message: &str) -> String {
    return format!("[WARNING]: {}", message);
}

pub fn error(message: &str) -> String {
    return format!("[ERROR]: {}", message);
}

pub fn debug(message: &str) -> String {
    return format!("[DEBUG]: {}", message);
}