package human.resource.mgmt.domain;

import human.resource.mgmt.ScheduleApplication;
import human.resource.mgmt.domain.CalendarRegistered;
import java.util.Date;
import java.util.List;
import javax.persistence.*;
import lombok.Data;

@Entity
@Table(name = "Calendar_table")
@Data
public class Calendar {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private String userId;

    @OneToMany
    private List<Event> events;

    @PostPersist
    public void onPostPersist() {
        CalendarRegistered calendarRegistered = new CalendarRegistered(this);
        calendarRegistered.publishAfterCommit();
    }

    public static CalendarRepository repository() {
        CalendarRepository calendarRepository = ScheduleApplication.applicationContext.getBean(
            CalendarRepository.class
        );
        return calendarRepository;
    }

    public void addCalendar(AddCalendarCommand addCalendarCommand) {
        ScheduleAdded scheduleAdded = new ScheduleAdded(this);
        scheduleAdded.publishAfterCommit();
    }

    public void cancelCalendar(CancelCalendarCommand cancelCalendarCommand) {
        ScheduleCanceled scheduleCanceled = new ScheduleCanceled(this);
        scheduleCanceled.publishAfterCommit();
    }

    public void delaySchedule(DelayScheduleCommand delayScheduleCommand) {}
}
