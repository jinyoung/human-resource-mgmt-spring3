package human.resource.mgmt.infra;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import human.resource.mgmt.config.kafka.KafkaProcessor;
import human.resource.mgmt.domain.*;
import javax.naming.NameParser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.stream.annotation.StreamListener;
import org.springframework.kafka.support.Acknowledgment;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Transactional(rollbackFor = Exception.class)
public class PolicyHandler {

    @Autowired
    CalendarRepository calendarRepository;

    @StreamListener(KafkaProcessor.INPUT)
    public void whatever(
        @Payload String eventString,
        @Header(KafkaHeaders.ACKNOWLEDGMENT) Acknowledgment acknowledgment,
        @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) byte[] messageKey
    ) {
        /*
          // Call port method with received messageKey to publish msg to the same partition. //
          DomainClass.portMethod(eventString, new String(messageKey));
          
          // ,or //
          new EventRaised(domain Obj).publishAfterCommit(new String(messageKey));
          // manual Offset Commit. //
          acknowledgment.acknowledge();  
          */
    }

    @StreamListener(
        value = KafkaProcessor.INPUT,
        condition = "headers['type']=='VacationRegistered'"
    )
    public void wheneverVacationRegistered_AddCalendar(
        @Payload VacationRegistered vacationRegistered,
        @Header(KafkaHeaders.ACKNOWLEDGMENT) Acknowledgment acknowledgment,
        @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) byte[] messageKey
    ) {
        VacationRegistered event = vacationRegistered;
        System.out.println(
            "\n\n##### listener AddCalendar : " + vacationRegistered + "\n\n"
        );

        AddCalendarCommand addCalendarCommand = new AddCalendarCommand();

        calendarRepository
            .findById(event.getId())
            .ifPresent(calendar -> {
                calendar.addCalendar(addCalendarCommand);
            });

        // Manual Offset Commit //
        acknowledgment.acknowledge();
    }

    @StreamListener(
        value = KafkaProcessor.INPUT,
        condition = "headers['type']=='VacationCancelled'"
    )
    public void wheneverVacationCancelled_CancelCalendar(
        @Payload VacationCancelled vacationCancelled,
        @Header(KafkaHeaders.ACKNOWLEDGMENT) Acknowledgment acknowledgment,
        @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) byte[] messageKey
    ) {
        VacationCancelled event = vacationCancelled;
        System.out.println(
            "\n\n##### listener CancelCalendar : " + vacationCancelled + "\n\n"
        );

        CancelCalendarCommand cancelCalendarCommand = new CancelCalendarCommand();

        calendarRepository
            .findById(event.getId())
            .ifPresent(calendar -> {
                calendar.cancelCalendar(cancelCalendarCommand);
            });

        // Manual Offset Commit //
        acknowledgment.acknowledge();
    }

    @StreamListener(
        value = KafkaProcessor.INPUT,
        condition = "headers['type']=='VacationRejected'"
    )
    public void wheneverVacationRejected_CancelCalendar(
        @Payload VacationRejected vacationRejected,
        @Header(KafkaHeaders.ACKNOWLEDGMENT) Acknowledgment acknowledgment,
        @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) byte[] messageKey
    ) {
        VacationRejected event = vacationRejected;
        System.out.println(
            "\n\n##### listener CancelCalendar : " + vacationRejected + "\n\n"
        );

        CancelCalendarCommand cancelCalendarCommand = new CancelCalendarCommand();

        calendarRepository
            .findById(event.getId())
            .ifPresent(calendar -> {
                calendar.cancelCalendar(cancelCalendarCommand);
            });

        // Manual Offset Commit //
        acknowledgment.acknowledge();
    }

    @StreamListener(
        value = KafkaProcessor.INPUT,
        condition = "headers['type']=='EmployeeJoined'"
    )
    public void wheneverEmployeeJoined_RegisterCalendar(
        @Payload EmployeeJoined employeeJoined,
        @Header(KafkaHeaders.ACKNOWLEDGMENT) Acknowledgment acknowledgment,
        @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) byte[] messageKey
    ) {
        EmployeeJoined event = employeeJoined;
        System.out.println(
            "\n\n##### listener RegisterCalendar : " + employeeJoined + "\n\n"
        );

        Calendar calendar = new Calendar();
        calendarRepository.save(calendar);

        // Manual Offset Commit //
        acknowledgment.acknowledge();
    }
}
