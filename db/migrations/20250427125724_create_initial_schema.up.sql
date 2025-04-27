-- Create Appointments Table
CREATE TABLE appointments (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL
);

-- Create Attendees Table
CREATE TABLE attendees (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255),
    metadata TEXT
);

-- Create Appointments_Attendees Join Table
CREATE TABLE appointments_attendees (
    appointment_id VARCHAR(255) NOT NULL,
    attendee_id VARCHAR(255) NOT NULL,
    role VARCHAR(50),
    rsvp_status VARCHAR(50) NOT NULL,
    PRIMARY KEY (appointment_id, attendee_id, rsvp_status),
    INDEX idx_appointment_attendee (appointment_id, attendee_id),
    INDEX idx_rsvp_status (rsvp_status),
    CONSTRAINT fk_appointment FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE CASCADE,
    CONSTRAINT fk_attendee FOREIGN KEY (attendee_id) REFERENCES attendees(id) ON DELETE CASCADE
);