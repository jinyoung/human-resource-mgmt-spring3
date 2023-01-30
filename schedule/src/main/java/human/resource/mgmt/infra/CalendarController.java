package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import java.util.Optional;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
// @RequestMapping(value="/calendars")
@Transactional
public class CalendarController {

    @Autowired
    CalendarRepository calendarRepository;

    @RequestMapping(
        value = "calendars/{id}/add",
        method = RequestMethod.PUT,
        produces = "application/json;charset=UTF-8"
    )
    public Calendar addCalendar(
        @PathVariable(value = "id") String id,
        @RequestBody AddCalendarCommand addCalendarCommand,
        HttpServletRequest request,
        HttpServletResponse response
    ) throws Exception {
        System.out.println("##### /calendar/addCalendar  called #####");
        Optional<Calendar> optionalCalendar = calendarRepository.findById(id);

        optionalCalendar.orElseThrow(() -> new Exception("No Entity Found"));
        Calendar calendar = optionalCalendar.get();
        calendar.addCalendar(addCalendarCommand);

        calendarRepository.save(calendar);
        return calendar;
    }

    @RequestMapping(
        value = "calendars/{id}/cancel",
        method = RequestMethod.PUT,
        produces = "application/json;charset=UTF-8"
    )
    public Calendar cancelCalendar(
        @PathVariable(value = "id") String id,
        @RequestBody CancelCalendarCommand cancelCalendarCommand,
        HttpServletRequest request,
        HttpServletResponse response
    ) throws Exception {
        System.out.println("##### /calendar/cancelCalendar  called #####");
        Optional<Calendar> optionalCalendar = calendarRepository.findById(id);

        optionalCalendar.orElseThrow(() -> new Exception("No Entity Found"));
        Calendar calendar = optionalCalendar.get();
        calendar.cancelCalendar(cancelCalendarCommand);

        calendarRepository.save(calendar);
        return calendar;
    }

    @RequestMapping(
        value = "calendars/{id}/delayschedule",
        method = RequestMethod.PUT,
        produces = "application/json;charset=UTF-8"
    )
    public Calendar delaySchedule(
        @PathVariable(value = "id") String id,
        @RequestBody DelayScheduleCommand delayScheduleCommand,
        HttpServletRequest request,
        HttpServletResponse response
    ) throws Exception {
        System.out.println("##### /calendar/delaySchedule  called #####");
        Optional<Calendar> optionalCalendar = calendarRepository.findById(id);

        optionalCalendar.orElseThrow(() -> new Exception("No Entity Found"));
        Calendar calendar = optionalCalendar.get();
        calendar.delaySchedule(delayScheduleCommand);

        calendarRepository.save(calendar);
        return calendar;
    }
}
